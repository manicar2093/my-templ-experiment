package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manicar2093/echoroutesview"
	"templ-demo/cmd/service/controllers"
	"templ-demo/cmd/service/translations"
	"templ-demo/core"
	"templ-demo/core/apperrors"
	"templ-demo/core/connections"
	"templ-demo/core/converters"
	"templ-demo/core/logger"
	"templ-demo/core/validator"
	"templ-demo/internal/users"
	"templ-demo/pkg/config"
)

func main() {
	var (
		e                = echo.New()
		restBaseEndpoint = "/api/v1"
		webBaseEndpoint  = "/app"
		restBaseGroup    = e.Group(restBaseEndpoint)
		webBaseGroup     = e.Group(webBaseEndpoint)
		conf             = converters.Must(core.ParseConfig[config.Config]())
		dbConn           = connections.GetGormConnection(conf.DatabaseConnectionConfig)
	)
	logger.Config()

	// Use only for web app. echo do not allow to register this on a group :/
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))

	e.Use(core.I18nMiddleware(translations.Embed, "en"))

	webBaseGroup.Static("/assets", "./cmd/service/assets")
	webBaseGroup.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:X-XSRF-TOKEN",
	}))
	webBaseGroup.Use(core.SessionSecretKeyMiddleware(conf.SessionSecretKeyConfig))

	e.Use(middleware.Logger())

	userRepository := users.NewUserRepository(dbConn)
	core.RegisterController(webBaseGroup, controllers.NewUserWebController(
		userRepository,
	))
	core.RegisterController(restBaseGroup, controllers.NewUserController(
		userRepository,
	))
	if err := echoroutesview.RegisterRoutesViewer(e); err != nil {
		panic(err)
	}

	e.HTTPErrorHandler = apperrors.HandlerWEcho
	e.Validator = validator.NoValidatorWarning{}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}

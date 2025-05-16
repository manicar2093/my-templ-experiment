package main

import (
	"fmt"
	"templ-demo/cmd/web/controllers"
	"templ-demo/core"
	"templ-demo/core/apperrors"
	"templ-demo/core/connections"
	"templ-demo/core/converters"
	"templ-demo/core/logger"
	"templ-demo/core/validator"
	"templ-demo/internal/user"
	"templ-demo/pkg/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/manicar2093/echoroutesview"
)

func main() {
	var (
		e            = echo.New()
		baseEndpoint = ""
		baseGroup    = e.Group(baseEndpoint)
		conf         = converters.Must(core.ParseConfig[config.Config]())
		dbConn       = connections.GetGormConnection(conf.DatabaseConnectionConfig)
	)
	logger.Config()
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Use(middleware.Logger())
	core.RegisterController(baseGroup, controllers.NewUserController(
		user.NewUserRepository(dbConn),
	))
	echoroutesview.RegisterRoutesViewer(e)
	e.Static("/assets", "./cmd/web/assets")

	e.HTTPErrorHandler = apperrors.HandlerWEcho
	e.Validator = validator.NewGooKitValidator()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}

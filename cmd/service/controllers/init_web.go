package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"templ-demo/cmd/service/controllers/initpages"
	"templ-demo/core"
)

type InitWebController struct{}

func NewInitWebController() *InitWebController {
	return &InitWebController{}
}

func (c *InitWebController) SetUpRoutes(group *echo.Group) {
	group.GET("/initial", c.GetHandler)
}

func (c *InitWebController) GetHandler(ctx echo.Context) error {
	return core.Render(ctx, http.StatusOK, initpages.HomePage())
}

package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type InitRestController struct{}

func NewInitRestController() *InitRestController {
	return &InitRestController{}
}

func (c *InitRestController) SetUpRoutes(group *echo.Group) {
	group.GET("/initial", c.GetHandler)
}

func (c *InitRestController) GetHandler(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello from your new API :D")
}

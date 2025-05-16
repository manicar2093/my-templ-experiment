package core

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func BindAndValidate(ctx echo.Context, i any) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}
	if err := ctx.Validate(i); err != nil {
		return err
	}
	return nil
}

func GetFromEchoContext[T any](ctx echo.Context, key string) T {
	return ctx.Get(key).(T)
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

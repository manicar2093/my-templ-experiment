package core

import (
	"context"
	"time"

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

type echoContextTemplWrapper struct {
	echo.Context
}

func (e echoContextTemplWrapper) Deadline() (deadline time.Time, ok bool) {
	return e.Context.Request().Context().Deadline()
}

func (e echoContextTemplWrapper) Done() <-chan struct{} {
	return e.Request().Context().Done()
}

func (e echoContextTemplWrapper) Err() error {
	return e.Context.Request().Context().Err()
}

func (e echoContextTemplWrapper) Value(key any) any {
	keyString, keyIsString := key.(string)
	if keyIsString {
		return e.Get(keyString)
	}

	return e.Request().Context().Value(key)
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(echoContextTemplWrapper{ctx}, buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func GetCqrfToken(ctx context.Context) string {
	return ctx.Value("csrf").(string)
}

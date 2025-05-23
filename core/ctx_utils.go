package core

import (
	"github.com/labstack/echo/v4"
)

// SessionSecretKey is the key used to store the session_key in the context. This is used to encrypt the session cookie.

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

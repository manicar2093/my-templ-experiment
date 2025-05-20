package core

import "github.com/labstack/echo/v4"

type SessionSecretKeyConfig struct {
	SessionSecretKey string `env:"SESSION_SECRET_KEY" validate:"required"`
}

func SessionSecretKeyMiddleware(sessionSecret SessionSecretKeyConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(SessionSecretKey, sessionSecret)
			return next(c)
		}
	}
}

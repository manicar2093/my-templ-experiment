package core

import (
	"context"
	"encoding/gob"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"templ-demo/cmd/web/ui/components/toast"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const (
	EchoContextTemplWrapperKey = "echoCtx"
	// SessionSecretKey is the key used to store the session_key in the context. This is used to encrypt the session cookie.
	SessionSecretKey = "session_key"
	sessionName      = "flash_messages"
)

type FlashMessage struct {
	Variant toast.Variant
	Message string
	Title   string
}

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

type EchoContextTemplWrapper struct {
	echo.Context
}

func (c *EchoContextTemplWrapper) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Request().Context().Deadline()
}

func (c *EchoContextTemplWrapper) Done() <-chan struct{} {
	return c.Request().Context().Done()
}

func (c *EchoContextTemplWrapper) Err() error {
	return c.Context.Request().Context().Err()
}

func (c *EchoContextTemplWrapper) Value(key any) any {
	return c.Request().Context().Value(key)
}

func (c *EchoContextTemplWrapper) GetCqrfToken() string {
	return c.Context.Get("csrf").(string)
}

func (c *EchoContextTemplWrapper) GetFlash() []FlashMessage {
	session, _ := getCookieStore(c.Context).Get(c.Request(), sessionName)
	fm := session.Flashes()
	// If we have some messages.
	if len(fm) > 0 {
		session.Save(c.Request(), c.Response())
		// Initiate a strings slice to return messages.
		var flashes []FlashMessage
		for _, fl := range fm {
			flashes = append(flashes, fl.(FlashMessage))
		}

		return flashes
	}
	return nil
}

func SetFlash(ctx echo.Context, flash FlashMessage) {
	session, _ := getCookieStore(ctx).Get(ctx.Request(), sessionName)
	gob.Register(FlashMessage{})
	session.AddFlash(flash)

	if err := session.Save(ctx.Request(), ctx.Response()); err != nil {
		logrus.Error(err)
	}
}

func getCookieStore(ctx echo.Context) *sessions.CookieStore {
	// In real-world applications, use env variables to store the session key.
	sessionKey := ctx.Get(SessionSecretKey).(SessionSecretKeyConfig)
	return sessions.NewCookieStore([]byte(sessionKey.SessionSecretKey))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	renderContext := context.WithValue(ctx.Request().Context(), EchoContextTemplWrapperKey, &EchoContextTemplWrapper{ctx})

	if err := t.Render(renderContext, buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func GetEchoTemplContext(ctx context.Context) *EchoContextTemplWrapper {
	return ctx.Value(EchoContextTemplWrapperKey).(*EchoContextTemplWrapper)
}

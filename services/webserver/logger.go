package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ungame/go-signup/logext"
)

type Logger struct {
	logger logext.Logger
}

func NewLogger(logger logext.Logger) Handler {
	return &Logger{
		logger: logger,
	}
}

func (l Logger) Register(app *fiber.App) {
	app.Use(func(ctx *fiber.Ctx) error {
		l.logger.Debug("%s", ctx.Request().String())
		err := ctx.Next()
		if err != nil {
			l.logger.Error("%s", err.Error())
		}
		l.logger.Debug("%s", ctx.Response().String())
		return err
	})
}

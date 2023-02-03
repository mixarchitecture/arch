package http

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/mixarchitecture/arch/auth/app"
	"github.com/mixarchitecture/arch/auth/delivery/http/mapper"
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/mixarchitecture/arch/shared/validator"
)

type Server struct {
	app       app.Application
	i18n      i18n.I18n
	validator validator.Validator
	ctx       context.Context
	mapper    *mapper.Mapper
}

type Config struct {
	App       app.Application
	I18n      i18n.I18n
	Validator validator.Validator
	Context   context.Context
}

func New(config Config) Server {
	return Server{
		app:       config.App,
		i18n:      config.I18n,
		validator: config.Validator,
		ctx:       config.Context,
		mapper:    mapper.New(),
	}
}

func (h Server) Load(router fiber.Router) fiber.Router {
	router.Post("/register", h.wrapWithTimeout(h.Register))
	router.Post("/login", h.wrapWithTimeout(h.Login))
	router.Post("/logout", h.wrapWithTimeout(h.Logout))
	router.Put("/extend-token", h.wrapWithTimeout(h.RefreshToken))
	return router
}

/*
func (h Server) parseBody(c *fiber.Ctx, d interface{}) {
	parser.ParseBody(c, h.validator, h.i18n, d)
}

func (h Server) parseParams(c *fiber.Ctx, d interface{}) {
	parser.ParseParams(c, h.validator, h.i18n, d)
}

func (h Server) parseQuery(c *fiber.Ctx, d interface{}) {
	parser.ParseQuery(c, h.validator, h.i18n, d)
}
*/

func (h Server) wrapWithTimeout(fn fiber.Handler) fiber.Handler {
	return timeout.New(fn, 10*time.Second)
}

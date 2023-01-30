package http

import (
	"github.com/mixarchitecture/arch/shared/i18n"
	httpI18n "github.com/mixarchitecture/arch/shared/server/http/i18n"
	"github.com/mixarchitecture/arch/shared/server/http/result"

	"github.com/mixarchitecture/arch/example/src/delivery/http/dtos"

	"github.com/gofiber/fiber/v2"
)

func (s Server) GetExample(ctx *fiber.Ctx) error {
	req := &dtos.GetExampleRequest{}
	s.parseParams(ctx, req)
	l, a := httpI18n.GetLanguagesInContext(s.i18n, ctx)
	res, err := s.app.Queries.GetExample.Handle(s.ctx, s.mapper.Example.ReqToGetExample(req))
	if err != nil {
		return s.parseErr(err, l, a)
	}
	msg := s.i18n.Translate(Messages.Success.GetExample, l, a)
	return result.SuccessDetail(msg, s.mapper.Example.GetExampleQueryToRes(&res))
}

func (s Server) ListExample(ctx *fiber.Ctx) error {
	req := &dtos.ListExampleRequest{}
	s.parseQuery(ctx, req)
	req.Default()
	l, a := httpI18n.GetLanguagesInContext(s.i18n, ctx)
	res, err := s.app.Queries.ListExample.Handle(s.ctx, s.mapper.Example.ReqToListExample(req))
	if err != nil {
		return s.parseErr(err, l, a)
	}
	msg := s.i18n.Translate(Messages.Success.ListExample, l, a)
	return result.SuccessDetail(msg, s.mapper.Example.ListExampleQueryToRes(&res))
}

func (s Server) CreateExample(ctx *fiber.Ctx) error {
	req := &dtos.CreateExampleRequest{}
	s.parseBody(ctx, req)
	l, a := httpI18n.GetLanguagesInContext(s.i18n, ctx)
	_, err := s.app.Commands.CreateExample.Handle(s.ctx, s.mapper.Example.ReqToCreateExample(req))
	if err != nil {
		return s.parseErr(err, l, a)
	}
	msg := s.i18n.Translate(Messages.Success.CreateExample, l, a)
	return result.SuccessDetail(msg, req, fiber.StatusCreated)
}

func (s Server) UpdateExample(ctx *fiber.Ctx) error {
	req := &dtos.UpdateExampleRequest{}
	s.parseParams(ctx, req)
	s.parseBody(ctx, req)
	l, a := httpI18n.GetLanguagesInContext(s.i18n, ctx)
	_, err := s.app.Commands.UpdateExample.Handle(s.ctx, s.mapper.Example.ReqToUpdateExample(req))
	if err != nil {
		return s.parseErr(err, l, a)
	}
	msg := s.i18n.Translate(Messages.Success.UpdateExample, l, a)
	return result.SuccessDetail(msg, req)
}

func (s Server) parseErr(err error, langs ...string) error {
	i, ok := err.(*i18n.I18nError)
	if ok {
		return result.Error(s.i18n.TranslateFromError(*i, langs...))
	}
	return result.Error(s.i18n.Translate(Messages.Error.Unexpected, langs...))
}

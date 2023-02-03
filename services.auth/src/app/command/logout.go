package command

import (
	"context"

	"github.com/mixarchitecture/arch/auth/config"
	"github.com/mixarchitecture/arch/auth/domain/user"
	"github.com/mixarchitecture/arch/shared/auth/token"
	"github.com/mixarchitecture/arch/shared/decorator"
	"github.com/mixarchitecture/arch/shared/events"
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/sirupsen/logrus"
)

type LogoutCommand struct {
	Token string
}

type LogoutResult struct{}

type LogoutHandler decorator.CommandHandler[LogoutCommand, *LogoutResult]

type logoutHandler struct {
	authTopics config.AuthTopics
	publisher  events.Publisher
	tokenSrv   token.Service
	errors     user.Errors
}

type LogoutHandlerConfig struct {
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	TokenSrv      token.Service
	Errors        user.Errors
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewLogoutHandler(config LogoutHandlerConfig) LogoutHandler {
	return decorator.ApplyCommandDecorators[LogoutCommand, *LogoutResult](
		logoutHandler{
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			tokenSrv:   config.TokenSrv,
			errors:     config.Errors,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h logoutHandler) Handle(ctx context.Context, command LogoutCommand) (*LogoutResult, *i18n.I18nError) {
	err := h.tokenSrv.Expire(command.Token)
	if err != nil {
		return nil, h.errors.Failed("logout")
	}
	return &LogoutResult{}, nil
}

package command

import (
	"context"

	"github.com/mixarchitecture/arch/auth/src/config"
	"github.com/mixarchitecture/arch/auth/src/domain/user"
	"github.com/mixarchitecture/arch/shared/auth/token"
	"github.com/mixarchitecture/arch/shared/decorator"
	"github.com/mixarchitecture/arch/shared/events"
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/mixarchitecture/arch/shared/jwt"
	"github.com/sirupsen/logrus"
)

type RefreshTokenCommand struct {
	Token string
	Claim *jwt.UserClaim
}

type RefreshTokenResult struct {
	Token string
}

type RefreshTokenHandler decorator.CommandHandler[RefreshTokenCommand, *RefreshTokenResult]

type refreshTokenHandler struct {
	authTopics config.AuthTopics
	publisher  events.Publisher
	tokenSrv   token.Service
	errors     user.Errors
}

type RefreshTokenHandlerConfig struct {
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	TokenSrv      token.Service
	Errors        user.Errors
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewRefreshTokenHandler(config RefreshTokenHandlerConfig) RefreshTokenHandler {
	return decorator.ApplyCommandDecorators[RefreshTokenCommand, *RefreshTokenResult](
		refreshTokenHandler{
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			tokenSrv:   config.TokenSrv,
			errors:     config.Errors,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h refreshTokenHandler) Handle(ctx context.Context, command RefreshTokenCommand) (*RefreshTokenResult, *i18n.I18nError) {
	err := h.tokenSrv.Expire(command.Token)
	if err != nil {
		return nil, h.errors.Failed("refresh token")
	}
	tkn, err := h.tokenSrv.Generate(command.Claim)
	if err != nil {
		return nil, h.errors.Failed("refresh token")
	}
	return &RefreshTokenResult{
		Token: tkn,
	}, nil
}

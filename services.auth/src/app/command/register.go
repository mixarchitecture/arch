package command

import (
	"context"

	"github.com/mixarchitecture/arch/auth/src/config"
	"github.com/mixarchitecture/arch/auth/src/domain/user"
	"github.com/mixarchitecture/arch/shared/auth/token"
	"github.com/mixarchitecture/arch/shared/cipher"
	"github.com/mixarchitecture/arch/shared/decorator"
	"github.com/mixarchitecture/arch/shared/events"
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/sirupsen/logrus"
)

type RegisterCommand struct {
	Email    string
	Password string
}

type RegisterResult struct {
	Token string
}

type RegisterHandler decorator.CommandHandler[RegisterCommand, *RegisterResult]

type registerHandler struct {
	userRepo   user.Repository
	authTopics config.AuthTopics
	publisher  events.Publisher
	errors     user.Errors
	tokenSrv   token.Service
}

type RegisterHandlerConfig struct {
	UserRepo      user.Repository
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	TokenSrv      token.Service
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
	Errors        user.Errors
}

func NewRegisterHandler(config RegisterHandlerConfig) RegisterHandler {
	return decorator.ApplyCommandDecorators[RegisterCommand, *RegisterResult](
		registerHandler{
			userRepo:   config.UserRepo,
			authTopics: config.AuthTopics,
			publisher:  config.Publisher,
			errors:     config.Errors,
			tokenSrv:   config.TokenSrv,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h registerHandler) Handle(ctx context.Context, command RegisterCommand) (*RegisterResult, *i18n.I18nError) {
	already, err := h.userRepo.GetByEmail(ctx, command.Email)
	if err != nil {
		return nil, err
	}
	if already != nil {
		return nil, h.errors.AlreadyExists(command.Email)
	}
	pw, error := cipher.Hash(command.Password)
	if error != nil {
		return nil, h.errors.Failed("hash")
	}
	user, err := h.userRepo.Create(ctx, command.Email, pw)
	if err != nil {
		return nil, err
	}
	tkn, error := h.tokenSrv.Generate(user.ToJwtClaims())
	if error != nil {
		return nil, h.errors.Failed("token")
	}
	_ = h.publisher.Publish(h.authTopics.Registered, user)
	return &RegisterResult{
		Token: tkn,
	}, err
}

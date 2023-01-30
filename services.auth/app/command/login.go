package command

import (
	"context"

	"github.com/mixarchitecture/arch/auth/config"
	"github.com/mixarchitecture/arch/auth/domain/user"
	"github.com/mixarchitecture/arch/shared/auth/token"
	"github.com/mixarchitecture/arch/shared/cipher"
	"github.com/mixarchitecture/arch/shared/decorator"
	"github.com/mixarchitecture/arch/shared/events"
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/sirupsen/logrus"
)

type LoginCommand struct {
	Email    string
	Password string
}

type LoginResult struct {
	Token string
}

type LoginHandler decorator.CommandHandler[LoginCommand, *LoginResult]

type loginHandler struct {
	userRepo   user.Repository
	authTopics config.AuthTopics
	publisher  events.Publisher
	errors     user.Errors
	tokenSrv   token.Service
}

type LoginHandlerConfig struct {
	UserRepo      user.Repository
	AuthTopics    config.AuthTopics
	Publisher     events.Publisher
	TokenSrv      token.Service
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
	Errors        user.Errors
}

func NewLoginHandler(config LoginHandlerConfig) LoginHandler {
	return decorator.ApplyCommandDecorators[LoginCommand, *LoginResult](
		loginHandler{
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

func (h loginHandler) Handle(ctx context.Context, command LoginCommand) (*LoginResult, *i18n.I18nError) {
	user, err := h.userRepo.GetByEmail(ctx, command.Email)
	if err != nil {
		return nil, err
	}
	if err := cipher.Compare(user.Password, command.Password); err != nil {
		_ = h.publisher.Publish(h.authTopics.LoginFailed, user)
		return nil, h.errors.InvalidPassword()
	}
	tkn, error := h.tokenSrv.Generate(user.ToJwtClaims())
	if error != nil {
		return nil, h.errors.Failed("token")
	}
	_ = h.publisher.Publish(h.authTopics.LoggedIn, user)
	return &LoginResult{
		Token: tkn,
	}, err
}

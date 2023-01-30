package claim_guard

import (
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/mixarchitecture/arch/shared/jwt"
	"github.com/mixarchitecture/arch/shared/server/http/auth/current_user"
	httpI18n "github.com/mixarchitecture/arch/shared/server/http/i18n"
	"github.com/mixarchitecture/arch/shared/server/http/result"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Claims []string
	I18n   *i18n.I18n
	MsgKey string
}

func New(cnf *Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		u := current_user.Parse(ctx)
		if checkClaims(u, cnf.Claims) {
			return ctx.Next()
		}
		l, a := httpI18n.GetLanguagesInContext(*cnf.I18n, ctx)
		return result.Error(cnf.I18n.Translate(cnf.MsgKey, l, a), fiber.StatusForbidden)
	}
}

func checkClaims(u *jwt.UserClaim, cs []string) bool {
	for _, c := range cs {
		if checkClaim(u, c) {
			return true
		}
	}
	return false
}

func checkClaim(u *jwt.UserClaim, c string) bool {
	for _, r := range u.Roles {
		if r == c {
			return true
		}
	}
	return false
}

package required_auth

import (
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/mixarchitecture/arch/shared/jwt"
	httpI18n "github.com/mixarchitecture/arch/shared/server/http/i18n"
	"github.com/mixarchitecture/arch/shared/server/http/result"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	I18n   i18n.I18n
	MsgKey string
}

func New(cnf Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := c.Locals("user")
		if u == nil || u.(*jwt.UserClaim).IsExpired() {
			l, a := httpI18n.GetLanguagesInContext(cnf.I18n, c)
			return result.Error(cnf.I18n.Translate(cnf.MsgKey, l, a), fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}

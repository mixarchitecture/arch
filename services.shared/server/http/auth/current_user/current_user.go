package current_user

import (
	"fmt"

	"github.com/mixarchitecture/arch/shared/auth/token"
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/mixarchitecture/arch/shared/jwt"
	"github.com/mixarchitecture/arch/shared/server/http/result"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	TokenSrv token.Service
	I18n     *i18n.I18n
	MsgKey   string
}

func New(cnf *Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := c.Cookies("token")
		if t == "" {
			t = getBearerToken(c)
			if t == "" {
				return c.Next()
			}
		}
		res, err := cnf.TokenSrv.Parse(t)
		if err != nil || res == nil {
			fmt.Printf("Error: %v", err)
			// l, a := cnf.I18n.GetLanguagesInContext(c)
			return result.Error(err.Error(), fiber.StatusUnauthorized) // cnf.I18n.Translate(cnf.MsgKey, l, a)
		}
		c.Locals("user", res)
		return c.Next()
	}
}

func getBearerToken(c *fiber.Ctx) string {
	b := c.Get("Authorization")
	if b == "" {
		return ""
	}
	return b[7:]
}

func Parse(c *fiber.Ctx) *jwt.UserClaim {
	return c.Locals("user").(*jwt.UserClaim)
}

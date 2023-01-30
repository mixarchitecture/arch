package i18n

import (
	i18nRoot "github.com/mixarchitecture/arch/shared/i18n"

	"github.com/gofiber/fiber/v2"
)

func GetLanguagesInContext(i i18nRoot.I18n, c *fiber.Ctx) (string, string) {
	l := c.Query("lang")
	a := c.Get("Accept-Language", i.Fallback)
	if l == "" {
		l = a
	}
	return l, a
}

func New(i i18nRoot.I18n) fiber.Handler {
	return func(c *fiber.Ctx) error {
		l, a := GetLanguagesInContext(i, c)
		c.Locals("lang", l)
		c.Locals("accept-language", a)
		return c.Next()
	}
}

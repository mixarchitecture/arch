package http

import (
	"fmt"

	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/mixarchitecture/arch/shared/server/http/error_handler"

	"github.com/goccy/go-json"

	i18nHttp "github.com/mixarchitecture/arch/shared/server/http/i18n"

	"github.com/mixarchitecture/arch/example/src/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host          string
	Port          int
	CreateHandler func(router fiber.Router) fiber.Router
	I18n          *i18n.I18n
	Cors          config.Cors
}

func RunServer(cfg Config) {
	addr := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	RunServerOnAddr(addr, cfg)
}

func RunServerOnAddr(addr string, cfg Config) {
	app := fiber.New(fiber.Config{
		Prefork: true,
		ErrorHandler: error_handler.New(error_handler.Config{
			// DfMsgKey: "error_internal_server_error",
			I18n: cfg.I18n,
		}),
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	group := app.Group("/api")
	setGlobalMiddlewares(app, cfg)
	cfg.CreateHandler(group)

	logrus.Infof("Starting server on %v", addr)
	if err := app.Listen(addr); err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func setGlobalMiddlewares(router fiber.Router, cfg Config) {
	router.Use(recover.New())
	router.Use(i18nHttp.New(*cfg.I18n))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Cors.AllowedOrigins,
		AllowHeaders:     cfg.Cors.AllowedHeaders,
		AllowMethods:     cfg.Cors.AllowedMethods,
		AllowCredentials: cfg.Cors.AllowCredentials,
	}))
	router.Use(compress.New(compress.Config{}))
}

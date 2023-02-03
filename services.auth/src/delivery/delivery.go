package delivery

import (
	"context"

	"github.com/mixarchitecture/arch/auth/app"
	"github.com/mixarchitecture/arch/auth/config"
	"github.com/mixarchitecture/arch/shared/events"
	"github.com/mixarchitecture/arch/shared/i18n"
)

type Delivery interface {
	Load()
}

type delivery struct {
	app         app.Application
	config      config.App
	i18n        *i18n.I18n
	ctx         context.Context
	eventEngine events.Engine
}

type Config struct {
	App         app.Application
	Config      config.App
	I18n        *i18n.I18n
	Ctx         context.Context
	EventEngine events.Engine
}

func New(config Config) Delivery {
	return &delivery{
		app:         config.App,
		config:      config.Config,
		i18n:        config.I18n,
		ctx:         config.Ctx,
		eventEngine: config.EventEngine,
	}
}

func (d *delivery) Load() {}

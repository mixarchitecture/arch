package main

import (
	"context"

	"github.com/mixarchitecture/arch/shared/env"
	"github.com/mixarchitecture/arch/shared/events/nats"
	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/mixarchitecture/arch/shared/logs"

	"github.com/mixarchitecture/arch/example/src/config"
	"github.com/mixarchitecture/arch/example/src/delivery"

	"github.com/mixarchitecture/arch/example/src/service"
)

func main() {
	logs.Init()
	ctx := context.Background()
	config := config.App{}
	env.Load(&config)
	i18n := i18n.New(config.I18n.Fallback)
	i18n.Load(config.I18n.Dir, config.I18n.Locales...)
	eventEngine := nats.New(nats.Config{
		Url:     config.Nats.Url,
		Streams: config.Nats.Streams,
	})
	app := service.NewApplication(config, eventEngine)
	delivery := delivery.New(delivery.Config{
		App:         app,
		Config:      config,
		I18n:        i18n,
		Ctx:         ctx,
		EventEngine: eventEngine,
	})
	delivery.Load()
}

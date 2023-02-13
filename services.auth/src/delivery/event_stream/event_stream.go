package event_stream

import (
	"github.com/mixarchitecture/arch/auth/src/app"
	"github.com/mixarchitecture/arch/auth/src/config"
	"github.com/mixarchitecture/arch/shared/events"
)

type Server struct {
	app    app.Application
	Topics config.AuthTopics
	engine events.Engine
}

type Config struct {
	App    app.Application
	Topics config.AuthTopics
	Engine events.Engine
}

func New(config Config) Server {
	return Server{
		app:    config.App,
		engine: config.Engine,
		Topics: config.Topics,
	}
}

func (s Server) Load() {
	// s.engine.Subscribe(s.Topics.UserUpdated, s.ListenUserUpdated)
}

package token

import (
	"github.com/mixarchitecture/arch/shared/db/redis"
	"github.com/mixarchitecture/arch/shared/events"
)

type App struct {
	Service    Service
	Handler    Handler
	Publisher  Publisher
	topic      string
	secretKey  string
	expiration int
	redis      redis.Service
	engine     events.Engine
}

type Config struct {
	SecretKey   string
	Expiration  int
	Redis       redis.Service
	EventEngine events.Engine
	Topic       string
}

func NewToken(cnf Config) *App {
	if cnf.Topic == "" {
		cnf.Topic = "Token"
	}
	srv := NewService(&SrvConfig{
		Secret:     cnf.SecretKey,
		Expiration: cnf.Expiration,
		Redis:      cnf.Redis,
	})
	return &App{
		Service:    srv,
		Handler:    NewHandler(HandlerConfig{Srv: srv, Engine: cnf.EventEngine, Topic: cnf.Topic}),
		Publisher:  NewPublisher(PublisherConfig{Engine: cnf.EventEngine, Topic: cnf.Topic}),
		topic:      cnf.Topic,
		secretKey:  cnf.SecretKey,
		expiration: cnf.Expiration,
		redis:      cnf.Redis,
		engine:     cnf.EventEngine,
	}
}

func (a *App) Init() {
	a.Handler.ListenAll()
}

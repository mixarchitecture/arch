package token

import "github.com/mixarchitecture/arch/shared/events"

type Publisher interface{}

type publisher struct {
	engine events.Engine
	topic  string
}

type PublisherConfig struct {
	Engine events.Engine
	Topic  string
}

func NewPublisher(cnf PublisherConfig) Publisher {
	if cnf.Topic == "" {
		cnf.Topic = "Token"
	}

	return &publisher{
		engine: cnf.Engine,
		topic:  cnf.Topic,
	}
}

func (p *publisher) createEventName(event string) string {
	return p.topic + "." + event
}

func (p *publisher) Created(token string) {
	_ = p.engine.Publish(p.createEventName("Created"), p.toDto(token))
}

func (p *publisher) Deleted(token string) {
	_ = p.engine.Publish(p.createEventName("Deleted"), p.toDto(token))
}

func (p *publisher) Extended(token string) {
	_ = p.engine.Publish(p.createEventName("Extended"), p.toDto(token))
}

func (p *publisher) toDto(token string) *dto {
	return &dto{
		Token: token,
	}
}

package nats

import (
	"encoding/json"

	"github.com/mixarchitecture/arch/shared/events"

	"github.com/nats-io/nats.go"
)

func (e *natsEngine) Open() error {
	for _, stream := range e.config.Streams {
		_, err := e.JS.AddStream(&nats.StreamConfig{
			Name:     stream,
			Subjects: []string{stream + ".*"},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *natsEngine) Publish(event string, data interface{}) error {
	p, err := e.Marshal(data)
	if err != nil {
		return err
	}
	return e.C.Publish(event, p)
}

func (e *natsEngine) Subscribe(event string, handler events.Handler) error {
	sub, err := e.C.Subscribe(event, func(msg *nats.Msg) {
		handler(msg.Data)
	})
	if err != nil {
		return err
	}
	e.subs = append(e.subs, sub)
	return nil
}

func (e *natsEngine) Unsubscribe(event string, handler events.Handler) error {
	for i, sub := range e.subs {
		if sub.Subject == event {
			err := sub.Unsubscribe()
			if err != nil {
				return err
			}
			e.subs = append(e.subs[:i], e.subs[i+1:]...)
		}
	}
	return nil
}

func (e *natsEngine) Marshal(data interface{}) ([]byte, error) {
	return json.Marshal(&data)
}

func (e *natsEngine) Unmarshal(data []byte, v interface{}) error {
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}

func (e *natsEngine) Close() error {
	for _, sub := range e.subs {
		err := sub.Unsubscribe()
		if err != nil {
			return err
		}
	}
	return nil
}

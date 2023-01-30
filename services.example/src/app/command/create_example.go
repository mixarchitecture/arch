package command

import (
	"context"

	"github.com/mixarchitecture/arch/shared/decorator"
	"github.com/mixarchitecture/arch/shared/events"
	"github.com/mixarchitecture/arch/shared/i18n"

	"github.com/mixarchitecture/arch/example/src/config"
	"github.com/mixarchitecture/arch/example/src/domain/example"

	"github.com/sirupsen/logrus"
)

type CreateExampleCommand struct {
	Field   string
	Content string
}

type CreateExampleHandler decorator.CommandHandler[CreateExampleCommand, *example.Example]

type createExampleHandler struct {
	exampleRepo   example.Repository
	exampleTopics config.ExampleTopics
	publisher     events.Publisher
}

type CreateExampleHandlerConfig struct {
	ExampleRepo   example.Repository
	ExampleTopics config.ExampleTopics
	Publisher     events.Publisher
	Logger        *logrus.Entry
	MetricsClient decorator.MetricsClient
}

func NewCreateExampleHandler(config CreateExampleHandlerConfig) CreateExampleHandler {
	return decorator.ApplyCommandDecorators[CreateExampleCommand, *example.Example](
		createExampleHandler{
			exampleRepo:   config.ExampleRepo,
			exampleTopics: config.ExampleTopics,
			publisher:     config.Publisher,
		},
		config.Logger,
		config.MetricsClient,
	)
}

func (h createExampleHandler) Handle(ctx context.Context, command CreateExampleCommand) (*example.Example, *i18n.I18nError) {
	example := &example.Example{
		Field:   command.Field,
		Content: command.Content,
	}
	exmp, err := h.exampleRepo.Create(ctx, example)
	if err != nil {
		return example, err
	}
	_ = h.publisher.Publish(h.exampleTopics.Created, example)
	return exmp, nil
}

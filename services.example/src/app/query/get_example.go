package query

import (
	"context"

	"github.com/mixarchitecture/arch/shared/decorator"
	"github.com/mixarchitecture/arch/shared/i18n"

	"github.com/mixarchitecture/arch/example/src/domain/example"

	"github.com/sirupsen/logrus"
)

type GetExampleQuery struct {
	Field string
}

type GetExampleResult struct {
	UUID    string
	Content string
	Field   string
}

type GetExampleHandler decorator.QueryHandler[GetExampleQuery, GetExampleResult]

type getExampleHandler struct {
	exampleRepo example.Repository
}

func NewGetExampleHandler(exampleRepo example.Repository, logger *logrus.Entry, metrics decorator.MetricsClient) GetExampleHandler {
	return decorator.ApplyQueryDecorators[GetExampleQuery, GetExampleResult](
		getExampleHandler{exampleRepo: exampleRepo},
		logger,
		metrics,
	)
}

func (h getExampleHandler) Handle(ctx context.Context, query GetExampleQuery) (GetExampleResult, *i18n.I18nError) {
	example, err := h.exampleRepo.Get(ctx, query.Field)
	if err != nil {
		return GetExampleResult{}, err
	}

	return GetExampleResult{Content: example.Content, Field: example.Field, UUID: example.UUID}, nil
}

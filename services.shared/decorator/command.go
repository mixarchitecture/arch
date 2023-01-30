package decorator

import (
	"context"
	"fmt"
	"strings"

	"github.com/mixarchitecture/arch/shared/i18n"
	"github.com/sirupsen/logrus"
)

type CommandHandler[C any, R any] interface {
	Handle(ctx context.Context, cmd C) (R, *i18n.I18nError)
}

func ApplyCommandDecorators[H any, R any](handler CommandHandler[H, R], logger *logrus.Entry, metricsClient MetricsClient) CommandHandler[H, R] {
	return &commandLoggingDecorator[H, R]{
		base: &commandMetricsDecorator[H, R]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

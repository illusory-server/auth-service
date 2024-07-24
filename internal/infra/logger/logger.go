package logger

import (
	"context"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/infra/tracing"
	"github.com/rs/zerolog"
	"io"
)

type logger struct {
	log *zerolog.Logger
}

func (l *logger) Error(ctx context.Context) *zerolog.Event {
	return l.log.Error().Ctx(ctx)
}

func (l *logger) Warn(ctx context.Context) *zerolog.Event {
	return l.log.Warn().Ctx(ctx)
}

func (l *logger) Info(ctx context.Context) *zerolog.Event {
	return l.log.Info().Ctx(ctx)
}

func (l *logger) Debug(ctx context.Context) *zerolog.Event {
	return l.log.Debug().Ctx(ctx)
}

func (l *logger) Fatal(ctx context.Context) *zerolog.Event {
	return l.log.Fatal().Ctx(ctx)
}

func (l *logger) Panic(ctx context.Context) *zerolog.Event {
	return l.log.Panic()
}

func newLogger(out io.Writer) domain.Logger {
	l := zerolog.New(out).With().Timestamp().Logger()
	l = l.Hook(requestIdHook{
		reqId: tracing.RequestIdKey,
	})
	return &logger{
		log: &l,
	}
}

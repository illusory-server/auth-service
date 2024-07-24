package domain

import (
	"context"
	"github.com/rs/zerolog"
)

type Logger interface {
	Error(context.Context) *zerolog.Event
	Warn(context.Context) *zerolog.Event
	Info(context.Context) *zerolog.Event
	Debug(context.Context) *zerolog.Event
	Fatal(context.Context) *zerolog.Event
	Panic(context.Context) *zerolog.Event
}

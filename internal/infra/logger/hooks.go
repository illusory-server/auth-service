package logger

import (
	"github.com/illusory-server/auth-service/internal/infra/tracing"
	"github.com/rs/zerolog"
)

type requestIdHook struct {
	reqId string
}

func (r requestIdHook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	ctx := e.GetCtx()
	requestId := tracing.GetRequestId(ctx)
	if requestId != "" {
		e.Str(r.reqId, requestId)
	}
}

package tracing

import (
	"context"
	"github.com/rs/xid"
	"google.golang.org/grpc/metadata"
)

const (
	RequestIdKey = "x-request-id"
)

func AddRequestId(ctx context.Context) context.Context {
	requestId := GetRequestId(ctx)
	if requestId == "" {
		requestId = xid.New().String()
	}
	return context.WithValue(ctx, RequestIdKey, requestId)
}

func AddRequestIdGrpc(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	var reqId string
	if !ok {
		reqId = xid.New().String()
	} else {

	}
	header, ok := md[RequestIdKey]
	if !ok || len(header) == 0 {
		reqId = xid.New().String()
	} else {
		reqId = header[0]
		if reqId == "" {
			reqId = xid.New().String()
		}
	}
	ctx = context.WithValue(ctx, RequestIdKey, reqId)
	return ctx
}

func GetRequestId(ctx context.Context) string {
	val, ok := ctx.Value(RequestIdKey).(string)
	if ok {
		return val
	}
	return ""
}

package interceptors

import (
	"context"
	"errors"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/illusory-server/auth-service/internal/domain"
	errorgrpc "github.com/illusory-server/auth-service/internal/transport/errors/error_grpc"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"google.golang.org/grpc"
)

func errorLogHandle(ctx context.Context, err *eerror.Error, log domain.Logger) {
	code := err.StatusCode
	isWarnStatusCode := code == eerror.ErrForbidden ||
		code == eerror.ErrNotFound ||
		code == eerror.ErrConflict ||
		code == eerror.ErrUnauthorized ||
		code == eerror.ErrBadRequest ||
		code == eerror.ErrDeadlineExceeded

	if isWarnStatusCode {
		l := log.Warn(ctx).
			Int("code", int(err.StatusCode))

		stack := eerror.GetStack(err)
		if stack != nil {
			l = l.Interface("stack_trace", stack)
		}

		cause := eerror.Cause(err)
		if cause != nil {
			l = l.Str("cause", cause.Error())
		}

		info := eerror.Info(err)
		if info != nil {
			l = l.Interface("info", info)
		}

		l.Msg(err.Message)
	}
}

func EerrorInterceptor(logger domain.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		result, err := handler(ctx, req)
		if err != nil {
			var e *eerror.Error
			err = eerror.Err(err).
				Stack(etrace.Func{
					Package:   "grpc",
					Name:      "grpc_handler",
					CauseFunc: info.FullMethod,
					CauseParams: etrace.FuncParams{
						"request": req,
					},
				}).
				Err()
			if !errors.As(err, &e) {
				logger.Warn(ctx).
					Err(err).
					Msg("error not implement to eerror library")
			} else {
				errorLogHandle(ctx, e, logger)
			}
			return nil, errorgrpc.Catch(err)
		}
		return result, nil
	}
}

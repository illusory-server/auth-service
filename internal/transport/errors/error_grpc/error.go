package errorgrpc

import (
	"errors"
	"github.com/OddEer0/Eer0/eerror"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Catch(err error) error {
	var eErr *eerror.Error
	ok := errors.As(err, &eErr)
	if !ok {
		return status.Error(codes.Internal, err.Error())
	}
	switch eErr.StatusCode {
	case eerror.ErrInternal:
		return status.Error(codes.Internal, err.Error())
	}
	return status.Error(codes.Internal, err.Error())
}

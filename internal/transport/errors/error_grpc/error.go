package errorgrpc

import (
	"github.com/illusory-server/auth-service/pkg/eerr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Catch(err error) error {
	code := eerr.Code(err)
	switch code {
	case eerr.ErrInternal:
		return status.Error(codes.Internal, err.Error())
	}
	return status.Error(codes.Internal, err.Error())
}

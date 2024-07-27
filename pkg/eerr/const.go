package eerr

const (
	NotCode               ErrorCode = 0
	ErrBadRequest         ErrorCode = 400
	ErrUnauthorized       ErrorCode = 401
	ErrForbidden          ErrorCode = 403
	ErrNotFound           ErrorCode = 404
	ErrDeadlineExceeded   ErrorCode = 408
	ErrConflict           ErrorCode = 409
	ErrUnprocessable      ErrorCode = 422
	ErrInternal           ErrorCode = 500
	ErrUnimplemented      ErrorCode = 501
	ErrBadGateway         ErrorCode = 502
	ErrServiceUnavailable ErrorCode = 503
	ErrGatewayTimeout     ErrorCode = 504

	MsgBadRequest         = "Bad request"
	MsgUnauthorized       = "Unauthorized"
	MsgForbidden          = "Forbidden"
	MsgNotFound           = "Not found"
	MsgDeadlineExceeded   = "Deadline exceeded"
	MsgConflict           = "Conflict"
	MsgUnprocessable      = "Unprocessable"
	MsgInternal           = "Internal"
	MsgUnimplemented      = "Unimplemented"
	MsgBadGateway         = "Bad Gateway"
	MsgServiceUnavailable = "Service Unavailable"
	MsgGatewayTimeout     = "Gateway Timeout"
)

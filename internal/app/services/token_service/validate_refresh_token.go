package tokenService

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/illusory-server/auth-service/internal/domain"
)

func (s *service) ValidateRefreshToken(ctx context.Context, refreshToken string) (*JwtUserData, error) {
	//cfg := s.cfg
	//token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(cfg.Secret.ApiKey), nil
	//})
	//if err != nil {
	//	var jwtErr *jwt.ValidationError
	//	if errors.As(err, &jwtErr) {
	//		return nil, jwtErrHandle(ctx, jwtErr, s.log)
	//	}
	//	s.log.ErrorContext(ctx, "validate jwt token error", slog.Any("cause", err), slog.String("token", refreshToken))
	//	return nil, domain.NewErr(domain.ErrInternalCode, domain.ErrInternalMessage)
	//}
	//if !token.Valid {
	//	s.log.ErrorContext(ctx, "invalid token", slog.Any("cause", err), slog.String("token", refreshToken))
	//	return nil, domain.NewErr(domain.ErrUnauthorizedCode, domain.ErrUnauthorizedMessage)
	//}
	//claim := token.Claims.(*CustomClaims)
	//return &claim.JwtUserData, nil
	return nil, nil
}

func jwtErrHandle(ctx context.Context, jwtErr *jwt.ValidationError, log domain.Logger) error {
	//if jwtErr.Errors&jwt.ValidationErrorMalformed != 0 {
	//	log.ErrorContext(ctx, "uncorrected jwt token")
	//} else if jwtErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	//	log.ErrorContext(ctx, "token does not work or time")
	//} else {
	//	log.ErrorContext(ctx, "token validate error")
	//}
	//return domain.NewErr(domain.ErrUnauthorizedCode, domain.ErrUnauthorizedMessage)
	return nil
}

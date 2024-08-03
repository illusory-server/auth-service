package tokenService

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/golang-jwt/jwt"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"github.com/pkg/errors"
)

func (s *service) ValidateRefreshToken(ctx context.Context, refreshToken string) (*JwtUserData, error) {
	cfg := s.cfg
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret.AccessApiKey), nil
	})
	if err != nil {
		var jwtErr *jwt.ValidationError
		tr := traceTokenService.OfName("ValidateRefreshToken").
			OfCauseMethod("jwt.ParseWithClaims").
			OfCauseParams(etrace.FuncParams{
				"refresh_token": refreshToken,
			})
		if errors.As(err, &jwtErr) {
			return nil, jwtErrHandle(ctx, jwtErr, s.log, refreshToken, tr)
		}
		return nil, eerror.Err(err).
			Code(eerror.ErrInternal).
			Msg(eerror.MsgInternal).
			Stack(tr).
			Err()
	}
	if !token.Valid {
		tr := traceTokenService.OfName("ValidateRefreshToken").
			OfCauseMethod("token.Valid").
			OfCauseParams(etrace.FuncParams{
				"refresh_token": refreshToken,
			})
		return nil, eerror.Err(err).Code(eerror.ErrUnauthorized).Msg(eerror.MsgUnauthorized).Stack(tr).Err()
	}
	claim, ok := token.Claims.(*CustomClaims)
	if !ok {
		s.log.Error(ctx).
			Str("refresh_token", refreshToken).
			Msg("invalid token")
		tr := traceTokenService.OfName("ValidateRefreshToken").
			OfCauseFunc("type assertion").
			OfCauseParams(etrace.FuncParams{
				"from_value": token.Claims,
				"to_type":    "CustomClaims",
			})
		return nil, eerror.New("invalid type assertion").
			Code(eerror.ErrInternal).
			Stack(tr).Err()
	}
	return &claim.JwtUserData, nil
}

func jwtErrHandle(ctx context.Context, jwtErr *jwt.ValidationError, log domain.Logger, token string, tr etrace.Method) error {
	eLog := log.Error(ctx).Str("refresh_token", token)
	if jwtErr.Errors&jwt.ValidationErrorMalformed != 0 {
		eLog.Msg("uncorrected jwt token")
	} else if jwtErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
		eLog.Msg("token does not work or time")
	} else {
		eLog.Msg("token is invalid")
	}
	return eerror.Err(jwtErr).
		Code(eerror.ErrUnauthorized).
		Msg(eerror.MsgUnauthorized).
		Stack(tr).
		Err()
}

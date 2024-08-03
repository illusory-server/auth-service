package tokenService

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/golang-jwt/jwt"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"github.com/pkg/errors"
)

func (s *service) ValidateAccessToken(ctx context.Context, accessToken string) (*JwtUserData, error) {
	cfg := s.cfg
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret.AccessApiKey), nil
	})
	if err != nil {
		var jwtErr *jwt.ValidationError
		tr := traceTokenService.OfName("ValidateAccessToken").
			OfCauseMethod("jwt.ParseWithClaims").
			OfCauseParams(etrace.FuncParams{
				"access_token": accessToken,
			})
		if errors.As(err, &jwtErr) {
			return nil, jwtErrHandle(ctx, jwtErr, s.log, accessToken, tr)
		}
		return nil, eerror.Err(err).Code(eerror.ErrInternal).Msg(eerror.MsgInternal).Stack(tr).Err()
	}
	if !token.Valid {
		tr := traceTokenService.OfName("ValidateAccessToken").
			OfCauseMethod("token.Valid").
			OfCauseParams(etrace.FuncParams{
				"access_token": accessToken,
			})
		return nil, eerror.Err(err).Code(eerror.ErrUnauthorized).Msg(eerror.MsgUnauthorized).Stack(tr).Err()
	}
	claim, ok := token.Claims.(*CustomClaims)
	if !ok {
		s.log.Error(ctx).
			Str("access_token", accessToken).
			Msg("invalid token")
		tr := traceTokenService.OfName("ValidateAccessToken").
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

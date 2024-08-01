package tokenService

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"github.com/pkg/errors"
)

func (s *service) ValidateRefreshToken(ctx context.Context, refreshToken string) (*JwtUserData, error) {
	cfg := s.cfg
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret.AccessApiKey), nil
	})
	if err != nil {
		var jwtErr *jwt.ValidationError
		if errors.As(err, &jwtErr) {
			return nil, eerr.Wrap(jwtErrHandle(ctx, jwtErr, s.log, refreshToken), "[service.ValidateRefreshToken] jwt.ParseWithClaims")
		}
		s.log.Error(ctx).
			Err(err).
			Str("refresh_token", refreshToken).
			Msg("parse refresh token failed")
		return nil, eerr.Wrap(err, "[service.ValidateRefreshToken] jwt.ParseWithClaims")
	}
	if !token.Valid {
		s.log.Error(ctx).
			Str("refresh_token", refreshToken).
			Msg("invalid token")
		return nil, eerr.New(eerr.ErrUnauthorized, eerr.MsgUnauthorized)
	}
	claim, ok := token.Claims.(*CustomClaims)
	if !ok {
		s.log.Error(ctx).
			Str("refresh_token", refreshToken).
			Msg("invalid token")
	}
	return &claim.JwtUserData, nil
}

func jwtErrHandle(ctx context.Context, jwtErr *jwt.ValidationError, log domain.Logger, token string) error {
	eLog := log.Error(ctx).Str("refresh_token", token)
	if jwtErr.Errors&jwt.ValidationErrorMalformed != 0 {
		eLog.Msg("uncorrected jwt token")
	} else if jwtErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
		eLog.Msg("token does not work or time")
	} else {
		eLog.Msg("token is invalid")
	}
	return eerr.New(eerr.ErrUnauthorized, eerr.MsgUnauthorized)
}

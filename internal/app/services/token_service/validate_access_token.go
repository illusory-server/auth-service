package tokenService

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"github.com/pkg/errors"
)

func (s *service) ValidateAccessToken(ctx context.Context, accessToken string) (*JwtUserData, error) {
	cfg := s.cfg
	token, err := jwt.ParseWithClaims(accessToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret.AccessApiKey), nil
	})
	if err != nil {
		var jwtErr *jwt.ValidationError
		if errors.As(err, &jwtErr) {
			return nil, eerr.Wrap(jwtErrHandle(ctx, jwtErr, s.log, accessToken), "[service.ValidateRefreshToken] jwt.ParseWithClaims")
		}
		s.log.Error(ctx).
			Err(err).
			Str("access_token", accessToken).
			Msg("parse refresh token failed")
		return nil, eerr.Wrap(err, "[service.ValidateRefreshToken] jwt.ParseWithClaims")
	}
	if !token.Valid {
		s.log.Error(ctx).
			Str("access_token", accessToken).
			Msg("invalid token")
		return nil, eerr.New(eerr.ErrUnauthorized, eerr.MsgUnauthorized)
	}
	claim, ok := token.Claims.(*CustomClaims)
	if !ok {
		s.log.Error(ctx).
			Str("access_token", accessToken).
			Msg("invalid token")
	}
	return &claim.JwtUserData, nil
}

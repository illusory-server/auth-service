package tokenService

import (
	"context"
	"github.com/golang-jwt/jwt"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/pkg/eerr"
	"time"
)

func (s *service) Generate(_ context.Context, data JwtUserData) (*appDto.JwtTokens, error) {
	cfg := s.cfg
	accessDuration, err := time.ParseDuration(cfg.Secret.AccessTokenTime)
	if err != nil {
		return nil, eerr.Wrap(err, "[service.Generate] time.ParseDuration(1)")
	}
	refreshDuration, err := time.ParseDuration(cfg.Secret.RefreshTokenTime)
	if err != nil {
		return nil, eerr.Wrap(err, "[service.Generate] time.ParseDuration(2)")
	}
	accessClaims := CustomClaims{
		JwtUserData:    data,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(accessDuration).Unix()},
	}
	refreshClaims := CustomClaims{
		JwtUserData:    data,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(refreshDuration).Unix()},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	accessTokenString, err := accessToken.SignedString([]byte(cfg.Secret.ApiKey))
	if err != nil {
		return nil, eerr.Wrap(err, "[service.Generate] jwt.SignedString(1)")
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.Secret.ApiKey))
	if err != nil {
		return nil, eerr.Wrap(err, "[service.Generate] jwt.SignedString(2)")
	}
	return &appDto.JwtTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

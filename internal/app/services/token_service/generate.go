package tokenService

import (
	"context"
	"github.com/OddEer0/Eer0/eerror"
	"github.com/golang-jwt/jwt"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/pkg/etrace"
	"time"
)

func (s *service) Generate(_ context.Context, data JwtUserData) (*appDto.JwtTokens, error) {
	cfg := s.cfg
	accessDuration, err := time.ParseDuration(cfg.Secret.AccessTokenTime)
	if err != nil {
		tr := traceTokenService.OfName("Generate").
			OfCauseMethod("time.ParseDuration(1)").
			OfCauseParams(etrace.FuncParams{
				"duration": cfg.Secret.AccessTokenTime,
			})
		return nil, eerror.Err(err).Stack(tr).Err()
	}
	refreshDuration, err := time.ParseDuration(cfg.Secret.RefreshTokenTime)
	if err != nil {
		tr := traceTokenService.OfName("Generate").
			OfCauseMethod("time.ParseDuration(2)").
			OfCauseParams(etrace.FuncParams{
				"duration": cfg.Secret.RefreshTokenTime,
			})
		return nil, eerror.Err(err).Stack(tr).Err()
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
	accessTokenString, err := accessToken.SignedString([]byte(cfg.Secret.AccessApiKey))
	if err != nil {
		tr := traceTokenService.OfName("Generate").
			OfCauseMethod("accessToken.SignedString(1)")
		return nil, eerror.Err(err).Stack(tr).Err()
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.Secret.RefreshApiKey))
	if err != nil {
		tr := traceTokenService.OfName("Generate").
			OfCauseMethod("refreshToken.SignedString(2)")
		return nil, eerror.Err(err).Stack(tr).Err()
	}
	return &appDto.JwtTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

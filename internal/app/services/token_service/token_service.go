package tokenService

import (
	"context"
	"github.com/golang-jwt/jwt"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/infra/config"
	"github.com/illusory-server/auth-service/pkg/etrace"
)

type (
	JwtUserData struct {
		Id   domain.Id `json:"id"`
		Role string    `json:"role"`
	}

	CustomClaims struct {
		JwtUserData `json:"jwtUserData"`
		jwt.StandardClaims
	}

	Service interface {
		Generate(ctx context.Context, data JwtUserData) (*appDto.JwtTokens, error)
		ValidateRefreshToken(ctx context.Context, refreshToken string) (*JwtUserData, error)
		ValidateAccessToken(ctx context.Context, accessToken string) (*JwtUserData, error)
	}

	service struct {
		log             domain.Logger
		cfg             *config.Config
		tokenRepository repository.TokenRepository
	}
)

var (
	traceTokenService = etrace.Method{
		Package: "tokenService",
		Type:    "tokenService",
	}
)

func New(logger domain.Logger, cfg *config.Config, tokenRepo repository.TokenRepository) Service {
	return &service{
		log:             logger,
		cfg:             cfg,
		tokenRepository: tokenRepo,
	}
}

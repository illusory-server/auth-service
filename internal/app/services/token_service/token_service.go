package tokenService

import (
	"context"
	"github.com/golang-jwt/jwt"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/model"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/infra/config"
)

type (
	JwtTokens struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	JwtUserData struct {
		Id   string `json:"id"`
		Role string `json:"role"`
	}

	CustomClaims struct {
		JwtUserData `json:"jwtUserData"`
		jwt.StandardClaims
	}

	Service interface {
		HasByValue(ctx context.Context, refreshToken string) (bool, error)
		Generate(ctx context.Context, data JwtUserData) (*JwtTokens, error)
		ValidateRefreshToken(ctx context.Context, refreshToken string) (*JwtUserData, error)
		Save(ctx context.Context, data appDto.SaveTokenServiceDto) (*model.Token, error)
		DeleteByValue(ctx context.Context, value string) error
	}

	service struct {
		log             domain.Logger
		cfg             *config.Config
		tokenRepository repository.TokenRepository
	}
)

func New(logger domain.Logger, cfg *config.Config, tokenRepo repository.TokenRepository) Service {
	return &service{
		log:             logger,
		cfg:             cfg,
		tokenRepository: tokenRepo,
	}
}

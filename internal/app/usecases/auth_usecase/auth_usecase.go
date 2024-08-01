package authUseCase

import (
	"context"
	appDto "github.com/illusory-server/auth-service/internal/app/app_dto"
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/domain/sql"
)

type (
	AuthResult struct {
		User   *appDto.PureUser
		Tokens *appDto.JwtTokens
	}

	UseCase interface {
		Registration(ctx context.Context, data *appDto.CreateUser) (*AuthResult, error)
		Login(ctx context.Context, data *appDto.LoginData) (*AuthResult, error)
		Logout(ctx context.Context, refreshToken string) error
		Refresh(ctx context.Context, refreshToken string) (*AuthResult, error)
	}

	useCase struct {
		log                domain.Logger
		userRepository     repository.UserRepository
		tokenService       tokenService.Service
		tokenRepository    repository.TokenRepository
		activateRepository repository.ActivateRepository
		banRepository      repository.BanRepository
		tx                 sql.Transactor
	}
)

func New(logger domain.Logger,
	userRepo repository.UserRepository,
	tokenServ tokenService.Service,
	tokenRepo repository.TokenRepository,
	activateRepository repository.ActivateRepository,
	banRepository repository.BanRepository,
	tx sql.Transactor) UseCase {
	return &useCase{
		log:                logger,
		userRepository:     userRepo,
		tokenService:       tokenServ,
		tokenRepository:    tokenRepo,
		activateRepository: activateRepository,
		banRepository:      banRepository,
		tx:                 tx,
	}
}

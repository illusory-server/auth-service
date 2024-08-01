package interactor

import (
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	authUseCase "github.com/illusory-server/auth-service/internal/app/usecases/auth_usecase"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/domain/sql"
	"github.com/illusory-server/auth-service/internal/infra/config"
	psqlActivateRepository "github.com/illusory-server/auth-service/internal/infra/storage/psql/psql_activate_repository"
	psqlBanRepository "github.com/illusory-server/auth-service/internal/infra/storage/psql/psql_ban_repository"
	psqlTokenRepository "github.com/illusory-server/auth-service/internal/infra/storage/psql/psql_token_repository"
	psqlUserRepository "github.com/illusory-server/auth-service/internal/infra/storage/psql/psql_user_repository"
)

type Dependencies struct {
	UserRepository     repository.UserRepository
	TokenRepository    repository.TokenRepository
	ActivateRepository repository.ActivateRepository
	TokenService       tokenService.Service
	AuthUseCase        authUseCase.UseCase
}

func New(cfg *config.Config, log domain.Logger, db sql.QueryExecutor, tx sql.Transactor, txController sql.TransactionController) *Dependencies {
	// postgres Repository initialize
	pgUserRepo := psqlUserRepository.New(db, log, txController)
	pgTokenRepo := psqlTokenRepository.New(db, log, txController)
	pgActivateRepo := psqlActivateRepository.New(db, log, txController)
	pgBanRepo := psqlBanRepository.New(log, txController, db)

	// app services initialize
	tokenServ := tokenService.New(log, cfg, pgTokenRepo)

	// app use case initialize
	authUCase := authUseCase.New(log, pgUserRepo, tokenServ, pgTokenRepo, pgActivateRepo, pgBanRepo, tx)

	return &Dependencies{
		UserRepository:     pgUserRepo,
		TokenRepository:    pgTokenRepo,
		ActivateRepository: pgActivateRepo,
		TokenService:       tokenServ,
		AuthUseCase:        authUCase,
	}
}

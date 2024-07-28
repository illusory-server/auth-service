package interactor

import (
	tokenService "github.com/illusory-server/auth-service/internal/app/services/token_service"
	userService "github.com/illusory-server/auth-service/internal/app/services/user_service"
	authUseCase "github.com/illusory-server/auth-service/internal/app/usecases/auth_usecase"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	"github.com/illusory-server/auth-service/internal/domain/sql"
	"github.com/illusory-server/auth-service/internal/infra/config"
	"github.com/illusory-server/auth-service/internal/infra/storage/psql"
	psqlTokenRepository "github.com/illusory-server/auth-service/internal/infra/storage/psql/psql_token_repository"
	psqlUserRepository "github.com/illusory-server/auth-service/internal/infra/storage/psql/psql_user_repository"
)

type Dependencies struct {
	UserRepository     repository.UserRepository
	TokenRepository    repository.TokenRepository
	ActivateRepository repository.ActivateRepository
	UserService        userService.Service
	TokenService       tokenService.Service
	AuthUseCase        authUseCase.UseCase
}

func New(cfg *config.Config, log domain.Logger, db sql.QueryExecutor) *Dependencies {
	// postgres Repository initialize
	txController := psql.PgxTransactionController{}
	pgUserRepo := psqlUserRepository.New(db, log, txController)
	pgTokenRepo := psqlTokenRepository.New(db, log, txController)
	//pgUserActivateRepo := postgresRepository.NewUserActivateRepository(log, db)
	var (
		pgActivateRepo repository.ActivateRepository
	)

	// app services initialize
	userServ := userService.New(log, pgUserRepo)
	tokenServ := tokenService.New(log, cfg, pgTokenRepo)

	// app use case initialize
	authUCase := authUseCase.New(log, userServ, pgUserRepo, tokenServ, pgTokenRepo)

	return &Dependencies{
		UserRepository:     pgUserRepo,
		TokenRepository:    pgTokenRepo,
		ActivateRepository: pgActivateRepo,
		UserService:        userServ,
		TokenService:       tokenServ,
		AuthUseCase:        authUCase,
	}
}

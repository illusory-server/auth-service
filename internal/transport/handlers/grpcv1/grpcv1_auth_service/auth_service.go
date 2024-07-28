package grpcv1AuthService

import (
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	authUseCase "github.com/illusory-server/auth-service/internal/app/usecases/auth_usecase"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/domain/repository"
	grpcMapper "github.com/illusory-server/auth-service/internal/transport/mapper/grpc_mapper"
)

var authMapper = grpcMapper.AuthMapper{}

type (
	Dependencies struct {
		UserRepository repository.UserRepository
		AuthUseCase    authUseCase.UseCase
		Log            domain.Logger
	}

	AuthServiceServer struct {
		authv1.UnimplementedAuthServiceServer
		authUseCase    authUseCase.UseCase
		userRepository repository.UserRepository
		log            domain.Logger
	}
)

func New(dependencies *Dependencies) authv1.AuthServiceServer {
	return &AuthServiceServer{authUseCase: dependencies.AuthUseCase, userRepository: dependencies.UserRepository, log: dependencies.Log}
}

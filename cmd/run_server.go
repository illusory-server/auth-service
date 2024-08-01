package cmd

import (
	"context"
	"github.com/illusory-server/auth-service/cmd/app"
	"github.com/illusory-server/auth-service/cmd/interactor"
	authv1 "github.com/illusory-server/auth-service/gen/auth"
	"github.com/illusory-server/auth-service/internal/infra/logger"
	"github.com/illusory-server/auth-service/internal/infra/storage/psql"
	grpcv1AuthService "github.com/illusory-server/auth-service/internal/transport/handlers/grpcv1/grpcv1_auth_service"
	"github.com/illusory-server/auth-service/internal/transport/interceptors"
	"google.golang.org/grpc"
	"net"
)

func RunServer(ctx context.Context, app *app.App, errCh chan<- error) {
	lis, err := net.Listen("tcp", app.Cfg.Server.Address)
	if err != nil {
		errCh <- err
		return
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptors.RequestMetricsInterceptor, psql.PingUnaryInterceptor(app.PSQL, app.Logger), logger.LoggingInterceptor(app.Logger)),
	)

	txController := psql.PgxTransactionController{}
	tx := &psql.PgxTransaction{Conn: app.PSQL, Log: app.Logger, TxController: txController}

	dependencies := interactor.New(app.Cfg, app.Logger, app.PSQL, tx, txController)
	authv1.RegisterAuthServiceServer(grpcServer, grpcv1AuthService.New(&grpcv1AuthService.Dependencies{
		UserRepository: dependencies.UserRepository,
		AuthUseCase:    dependencies.AuthUseCase,
		Log:            app.Logger,
	}))

	ch := make(chan error)
	go func() {
		app.Logger.Info(ctx).
			Str("address", app.Cfg.Server.Address).
			Msg("Server starting...")
		err := grpcServer.Serve(lis)
		if err != nil {
			app.Logger.Error(ctx).
				Err(err).
				Msg("grpc server error")
		}
		ch <- err
	}()
	select {
	case err := <-ch:
		app.Logger.Error(ctx).
			Err(err).
			Msg("Server listen error")
		errCh <- err
	case <-ctx.Done():
		app.Logger.Info(ctx).Msg("Server shutdown")
		return
	}
}

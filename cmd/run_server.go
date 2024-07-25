package cmd

import (
	"context"
	"github.com/illusory-server/auth-service/cmd/app"
	"net/http"
)

func RunServer(ctx context.Context, app *app.App, errCh chan<- error) {
	ch := make(chan error)
	go func() {
		app.Logger.Info(ctx).
			Str("port", "5500").
			Msg("Server starting...")
		err := http.ListenAndServe(":5500", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("work"))
		}))
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

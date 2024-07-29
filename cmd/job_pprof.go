package cmd

import (
	"context"
	"github.com/illusory-server/auth-service/cmd/app"
	"net/http"
	"net/http/pprof"
)

func JobPprof(ctx context.Context, app *app.App) {
	r := http.NewServeMux()
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":5050", r)
	if err != nil {
		panic(err)
	}
}

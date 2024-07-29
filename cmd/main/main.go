package main

import (
	"github.com/illusory-server/auth-service/cmd"
	"github.com/illusory-server/auth-service/cmd/app"
	"github.com/rs/zerolog/log"
	_ "net/http/pprof"
	"os"
	"runtime"
)

func main() {
	runtime.SetBlockProfileRate(1)
	application := app.Init()
	application.RegisterRunners(cmd.RunServer)
	application.RegisterJob(cmd.JobMetrics, &app.JobOptions{
		Name:  "prometheus_metrics",
		Retry: 3,
	})
	application.RegisterJob(cmd.JobPprof, &app.JobOptions{
		Name:  "pprof_debug",
		Retry: 3,
	})
	err := application.Run()

	if err != nil {
		log.Info().
			Err(err).
			Msg("app runner error")
		os.Exit(1)
		return
	}
}

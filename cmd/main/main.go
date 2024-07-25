package main

import (
	"github.com/illusory-server/auth-service/cmd"
	"github.com/illusory-server/auth-service/cmd/app"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	application := app.Init()
	application.RegisterRunners(cmd.RunServer)
	err := application.Run()

	if err != nil {
		log.Info().
			Err(err).
			Msg("app runner error")
		os.Exit(1)
		return
	}
}

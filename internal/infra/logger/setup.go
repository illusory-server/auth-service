package logger

import (
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/rs/zerolog"
	"io"
	"os"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
	EnvDev   = "dev"
	EnvTest  = "test"
)

type logOut struct {
	output    io.Writer
	errOutput io.Writer
}

func (l logOut) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}

func (l logOut) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level > zerolog.InfoLevel {
		return l.errOutput.Write(p)
	}
	return l.output.Write(p)
}

func MustLoad(env string) domain.Logger {
	switch env {
	case EnvDev:
		return setupDevLogger()
	case EnvTest:
		return setupDevLogger()
	case EnvLocal:
		return setupDevLogger()
	case EnvProd:
		return setupProdLogger()
	}
	panic("invalid env")
}

func setupDevLogger() domain.Logger {
	return newLogger(os.Stdout)
}

func setupProdLogger() domain.Logger {
	out := logOut{
		output:    os.Stdout,
		errOutput: os.Stderr,
	}
	return newLogger(out)
}

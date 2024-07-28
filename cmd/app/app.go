package app

import (
	"context"
	"errors"
	"github.com/illusory-server/auth-service/internal/domain"
	"github.com/illusory-server/auth-service/internal/infra/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	JobFunc    func(context.Context, *App)
	RunnerFunc func(ctx context.Context, app *App, errCh chan<- error)
	JobOptions struct {
		Retry int
		Name  string
	}
	job struct {
		name string
		fn   JobFunc
		opt  *JobOptions
	}
	App struct {
		Cfg     *config.Config
		Logger  domain.Logger
		PSQL    *pgxpool.Pool
		runners []RunnerFunc
		jobs    []job
	}
)

func (app *App) Run() (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	errCh := make(chan error, len(app.runners))
	for _, runner := range app.runners {
		go func(errCh chan<- error) {
			defer func() {
				if r := recover(); r != nil {
					errCh <- errors.New("panic")
				}
			}()
			runner(ctx, app, errCh)
		}(errCh)
	}
	app.jobsRun(ctx)
	err = <-errCh
	return err
}

func (app *App) RegisterRunners(runners ...RunnerFunc) {
	app.runners = append(app.runners, runners...)
}

func (app *App) RegisterJob(jobs JobFunc, opt *JobOptions) {
	j := job{
		fn:  jobs,
		opt: opt,
	}
	app.jobs = append(app.jobs, j)
}

func (app *App) jobsRun(ctx context.Context) {
	for _, j := range app.jobs {
		go func() {
			for i := 0; i < j.opt.Retry; i++ {
				ch := make(chan error)
				go func() {
					defer func() {
						if r := recover(); r != nil {
							ch <- errors.New("job: " + j.name + " panic")
						} else {
							ch <- nil
						}
					}()
					app.Logger.Info(ctx).
						Str("name", j.name).
						Msg("starting job...")
					j.fn(ctx, app)
				}()
				err := <-ch
				if err == nil {
					app.Logger.Info(ctx).
						Str("name", j.name).
						Msg("job success finished")
					break
				}
				app.Logger.Warn(ctx).
					Str("name", j.name).
					Err(err).
					Int("retry", i+1).
					Msg("job failed, retrying")
			}
		}()
	}
}

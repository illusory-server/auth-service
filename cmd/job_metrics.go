package cmd

import (
	"context"
	"github.com/illusory-server/auth-service/cmd/app"
	"github.com/illusory-server/auth-service/internal/infra/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func JobMetrics(ctx context.Context, app *app.App) {
	mux := http.NewServeMux()
	prometheus.MustRegister(metrics.RequestCount, metrics.RequestDuration)
	mux.Handle("/metrics", promhttp.Handler())

	app.Logger.Info(ctx).Msg("Metrics server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

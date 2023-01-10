package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/anzx/pkg/opentelemetry"
	"github.com/anzx/pkg/opentelemetry/exporters"
	"github.com/anzx/pkg/opentelemetry/metrics"
	"github.com/anzx/pkg/opentelemetry/trace"
)

func main() {
	//OTEL_EXPORTER_OTLP_ENDPOINT := "http://localhost:16686/"

	cfg := &opentelemetry.Config{
		Metrics: metrics.Config{
			Exporter: "stdout",
		},
		Trace: trace.Config{
			Exporter: "stdout",
		},
		Exporters: exporters.Exporters{
			Stdout: exporters.StdoutConfig{OTEL_EXPORTER_OTLP_ENDPOINT},
		},
	}

	ctx := context.Background()

	err := opentelemetry.Start(ctx, cfg)
	if err != nil {
		log.Fatalf("error starting opentelemetrty")
	}

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	app := NewApp(os.Stdin)
	go func() {
		if err := app.Run(ctx); err != nil {
			log.Fatalf("error running app: %s", err)
		}
	}()

	<-ctx.Done()
}

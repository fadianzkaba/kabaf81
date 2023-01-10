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
	var otelConfig *opentelemetry.Config

	// Set exporter apprilately depending on whether or not we detect the presence
	// of the OTEL_EXPORTER_OTLP_ENDPOINT environment variable (as per usptream OTEL docs).
	if endpoint, found := os.LookupEnv("OTEL_EXPORTER_OTLP_ENDPOINT"); found {
		otelConfig = &opentelemetry.Config{
			Metrics: metrics.Config{
				Exporter: "prometheus",
			},
			Trace: trace.Config{
				Exporter: "jaeger",
			},
			Exporters: exporters.Exporters{
				Jaeger: exporters.JaegerConfig{
					CollectorEndpoint: endpoint,
				},
			},
		}
	} else {
		otelConfig = &opentelemetry.Config{
			Metrics: metrics.Config{
				Exporter: "stdout",
			},
			Trace: trace.Config{
				Exporter: "stdout",
			},
			Exporters: exporters.Exporters{
				Stdout: exporters.StdoutConfig{},
			},
		}
	}

	ctx := context.Background()

	err := opentelemetry.Start(ctx, otelConfig)
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

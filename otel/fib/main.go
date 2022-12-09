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
	cfg := &opentelemetry.Config{
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

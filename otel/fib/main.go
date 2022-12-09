package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/anzx/pkg/opentelemetry"
	"github.com/anzx/pkg/opentelemetry/exporters"
	"github.com/anzx/pkg/opentelemetry/metrics"
	"github.com/anzx/pkg/opentelemetry/trace"
)

func main() {

	l := log.New(os.Stdout, "", 0)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}

	var ctx context.Context
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

	err := opentelemetry.Start(ctx, cfg)
	if err != nil {
		// handle startup error
	}

	_, spanEnd := opentelemetry.AddNamedSpan(context.Background(), "myservice", "testOtel")

	// this sleep simulates a work done in 500ms
	time.Sleep(time.Millisecond * 500)

	spanEnd()

	// This creates an int64 counter
	ctr := opentelemetry.CreateMetricInt64Counter("my.counter")

	// Record two data and provide a same label but different values.
	ctr.Add(context.Background(), 1, map[string]string{
		"my.label": "val1",
	})
	ctr.Add(context.Background(), 2, map[string]string{
		"my.label": "val1",
	})

	err = opentelemetry.Stop(ctx)
	if err != nil {
		// handle stop error
	}

	fmt.Println(ctr)

}

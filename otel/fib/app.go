package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/anzx/pkg/opentelemetry"
)

type App struct {
	r io.Reader
}

func NewApp(r io.Reader) *App {
	return &App{r: r}
}

func (a *App) Run(ctx context.Context) error {


	ctx, spanEnd := opentelemetry.AddSpan(ctx, "App")
	defer spanEnd()

	for {
		n, err := a.Poll(ctx)
		if err != nil {
			return err
		}

		a.Write(ctx, n)
	}
}

func (a *App) Poll(ctx context.Context) (uint, error) {
	_, spanEnd := opentelemetry.AddSpan(ctx, "App")
	defer spanEnd()

	log.Print("This what Fabicca would like to know")

	var n uint
	_, err := fmt.Fscanf(a.r, "%d\n", &n)
	return n, err
}

func (a *App) Write(ctx context.Context, n uint) {
	ctx, spanEnd := opentelemetry.AddSpan(ctx, "App")
	defer spanEnd()

	f, err := Fibonacci(ctx, n)
	if err != nil {
		log.Printf("Fibonacci(%d): %v\n", n, err)
	} else {
		log.Printf("Fibonacci(%d) = %d\n", n, f)
	}
}

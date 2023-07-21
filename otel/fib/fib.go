package main

import (
	"context"
	"time"

	"github.com/anzx/pkg/opentelemetry"
	
)

func Fibonacci(ctx context.Context, n uint) (uint64, error) {
	_, spanEnd := opentelemetry.AddSpan(ctx, "Main")
	time.Sleep(time.Microsecond * 500)

	defer spanEnd()

	if n <= 1 {
		return uint64(n), nil
	}

	var n2, n1 uint64 = 0, 1
	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1, nil
}

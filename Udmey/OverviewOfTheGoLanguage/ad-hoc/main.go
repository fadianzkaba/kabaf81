package main

import "context"

func main() {

	vat ctx context.Context()

	_, _ := program(ctx, "Fadi")

}

func program(ctx context.Context, name string) (string, string) {

	return "Hello", "error"

}

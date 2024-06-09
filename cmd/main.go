package main

import (
	"context"
	"fmt"

	"github.com/x-x-x-Ilya/accumulator-worker/internal/application"
)

const configPath = ".env"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := application.Start(ctx, configPath); err != nil {
		fmt.Println(err)
	}
}

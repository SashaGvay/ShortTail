package main

import (
	"context"
	"log"

	"short_tail/config"
	"short_tail/internal/root"
)

func main() {
	ctx := context.Background()

	conf, err := config.NewConfig()
	if err != nil {
		log.Panicf("config.NewConfig: %v", err)
	}

	app, err := root.New(ctx, conf)
	if err != nil {
		log.Panicf("root.New: %v", err)
	}

	if err = app.Run(ctx); err != nil {
		log.Panicf("app.Run: %v", err)
	}
}

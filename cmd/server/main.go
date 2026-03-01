package main

import (
	"context"
	"log"
	"time"

	"github.com/Ndbeloved/rate-limiter-go/internals/config"
	"github.com/Ndbeloved/rate-limiter-go/internals/server"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cfg := config.SetFlags()
	srv := server.New(cfg, ctx)
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}

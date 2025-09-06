package main

import (
	"context"
	"flag"
	"github.com/Ostmind/multiplicator/internal/servers/config"
	"github.com/Ostmind/multiplicator/internal/servers/logger"
	"github.com/Ostmind/multiplicator/internal/servers/server"
	"log"
)

func main() {
	rtp := flag.Float64("rtp", 0, "rtp parameter (must be > 0 and ≤ 1.0)")
	flag.Parse()

	if *rtp <= 0 || *rtp > 1.0 {
		log.Fatal("rtp должен быть > 0 и ≤ 1.0")
	}

	cfg := config.MustNew()
	sloger := logger.SetupLogger(cfg.EnvType)
	srv := server.New(sloger, *rtp)

	srv.Run(cfg.Host, cfg.Port)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServerShutdownTimeout)
	defer cancel()

	srv.Stop(ctx)
}

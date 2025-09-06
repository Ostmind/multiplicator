package main

import (
	"context"
	"flag"
	"github.com/Ostmind/multiplicator/internal/servers/config"
	"github.com/Ostmind/multiplicator/internal/servers/logger"
	"github.com/Ostmind/multiplicator/internal/servers/server"
	"log"
	"time"
)

func main() {
	rtpPtr := flag.Float64("rtp", 0, "rtp parameter (must be > 0 and ≤ 1.0)")
	flag.Parse()

	if *rtpPtr <= 0 || *rtpPtr > 1.0 {
		log.Fatal("rtp должен быть > 0 и ≤ 1.0")
	}

	cfg := config.MustNew()
	sloger := logger.SetupLogger(cfg.EnvType)
	srv := server.New(sloger, *rtpPtr)

	go srv.Run(cfg.Host, cfg.Port)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServerShutdownTimeout)
	defer cancel()

	time.Sleep(cfg.ServerShutdownTimeout)

	srv.Stop(ctx)
}

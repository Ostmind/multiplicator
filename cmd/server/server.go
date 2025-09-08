package main

import (
	"context"
	"flag"
	"github.com/Ostmind/multiplicator/internal/servers/config"
	"github.com/Ostmind/multiplicator/internal/servers/logger"
	"github.com/Ostmind/multiplicator/internal/servers/server"
	"log"
	"os"
	"os/signal"
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

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	sloger.Info("Received interrupt signal")
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServerShutdownTimeout)
	defer cancel()

	srv.Stop(ctx)
}

package server

import (
	"context"
	"github.com/Ostmind/multiplicator/internal/servers/config"
	"github.com/Ostmind/multiplicator/internal/servers/logger"
	"github.com/Ostmind/multiplicator/internal/servers/server"
	"os"
	"os/signal"
)

func StartServer(rtp float64) {

	cfg := config.MustNew()
	sloger := logger.SetupLogger(cfg.EnvType)
	srv := server.New(sloger, rtp)

	go srv.Run(cfg.Host, cfg.Port)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)

	<-stopChan
	sloger.Info("Received interrupt signal")
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServerShutdownTimeout)
	defer cancel()

	srv.Stop(ctx)
}

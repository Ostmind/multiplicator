package main

import (
	"context"
	"github.com/Ostmind/multiplicator/internal/servers/config"
	"github.com/Ostmind/multiplicator/internal/servers/logger"
	"github.com/Ostmind/multiplicator/internal/servers/server"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите параметр rtp при запуске: go run main.go <rtp>")
	}

	rtp, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil || rtp <= 0 || rtp > 1.0 {
		log.Fatal("rtp должен быть > 0 и ≤ 1.0")
	}

	cfg := config.MustNew()

	sloger := logger.SetupLogger(cfg.EnvType)

	srv := server.New(sloger, rtp)

	srv.Run(cfg.Host, cfg.Port)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.ServerShutdownTimeout)
	defer cancel()

	srv.Stop(ctx)
}

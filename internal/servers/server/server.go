package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	server *echo.Echo
	logger *slog.Logger
	rtp    float64
}

func New(logger *slog.Logger, rtp float64) *Server {
	server := echo.New()

	server.Use(logRequest(logger))

	numController := newNumberHandler(rtp, logger)

	server.GET("/get", numController.generateNumberHandler)

	return &Server{
		logger: logger,
		server: server,
		rtp:    rtp,
	}
}
func (s Server) Run(serverHost string, serverPort int) {
	s.logger.Info("Server is running on: ", "Host: ", serverHost, "Port: ", serverPort)

	if err := s.server.Start(fmt.Sprintf("%s:%d", serverHost, serverPort)); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			s.logger.Error("Server starting error: %v", slog.Any("error_details", err))
		}
	}
}

func (s Server) Stop(ctx context.Context) {
	s.logger.Info("Stopping server...")
	err := s.server.Shutdown(ctx)

	if err != nil {
		s.logger.Error("Error: ", slog.Any("error_details", err))
	}
}

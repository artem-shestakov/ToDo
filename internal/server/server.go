package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
	logger     *logrus.Logger
}

func NewServer(logger *logrus.Logger) *Server {
	return &Server{
		httpServer: &http.Server{},
		logger:     logger,
	}
}

func (s *Server) Run(address, port string, handler http.Handler) error {
	s.httpServer.Addr = address + ":" + port
	s.httpServer.Handler = handler
	s.logger.Info(fmt.Sprintf("Server starting on %s:%s...", address, port))
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

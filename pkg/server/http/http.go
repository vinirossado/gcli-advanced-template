package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"basic/pkg/logger"
)

type Server struct {
	*gin.Engine
	httpSrv *http.Server
	host    string
	port    int
	logger  *logger.Logger
}
type Option func(s *Server)

func NewServer(engine *gin.Engine, logger *logger.Logger, opts ...Option) *Server {
	s := &Server{
		Engine: engine,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}
func WithServerPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s,
	}

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Sugar().Fatalf("listen: %s\n", err)
	}

	return nil
}
func (s *Server) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	s.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		s.logger.Sugar().Fatal("Server forced to shutdown: ", err)
	}

	s.logger.Sugar().Info("Server exiting")
	return nil
}

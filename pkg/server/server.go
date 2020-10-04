package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rinchsan/txdb-todo/pkg/logger"
)

const (
	timeout = 10 * time.Second
)

func Run(r http.Handler, port int) int {
	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)
	s := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           r,
		ReadTimeout:       timeout,
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
	}

	go func() {
		logger.Infof("server listening on port %d", port)
		errCh <- s.ListenAndServe()
	}()

	select {
	case <-errCh:
		return 1
	case <-sigCh:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.Shutdown(ctx); err != nil {
			panic(err)
		}
		return 0
	}
}

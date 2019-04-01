package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) ServeHTTP() {
	g := s.errG

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	s.http = &http.Server{
		Addr:    s.cfg.HTTPListener(),
		Handler: tracing(nextRequestID)(logging(s.logger)(s.handler())),
	}

	g.Go(s.serveHTTP())
}

func (s *Server) shutdownHTTP(ctx context.Context) func() error {
	return func() error {
		s.log().Infof("Shutdown %s", s.http.Addr)

		return s.https.Shutdown(ctx)
	}
}

func (s *Server) serveHTTP() func() error {
	return func() error {
		var err error

		s.log().Infof("Listening on %s", s.http.Addr)

		if err = s.http.ListenAndServe(); err == http.ErrServerClosed {
			return nil
		}

		return err

	}
}

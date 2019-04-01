package server

import (
	"context"
	"time"

	"github.com/andersnormal/franz/config"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var _ Listener = (*Server)(nil)

func NewServer(ctx context.Context, cfg *config.Config) Listener {
	g, gtx := errgroup.WithContext(ctx)

	s := &Server{
		cfg:    cfg,
		errCtx: gtx,
		errG:   g,
	}

	return s
}

func (s *Server) Wait() error {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
		case <-s.errCtx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			g, gtx := errgroup.WithContext(ctx)

			if s.http != nil {
				g.Go(s.shutdownHTTP(gtx))
			}

			return g.Wait()
		}
	}
}

func (s *Server) config() *config.Config {
	return s.cfg
}

func (s *Server) log() *log.Entry {
	return s.logger
}

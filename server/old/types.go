package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/andersnormal/franz/config"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

// Listener describes the interface to a server
type Listener interface {
	// ServeHTTP is starting the HTTP listener
	ServeHTTP()
	// Wait is waiting for everything to end :-)
	Wait() error
}

type listener struct {
	// config to use with the server
	cfg *config.Config

	// logger attached to server
	logger *log.Entry

	// error Group
	errG *errgroup.Group

	// error Context
	errCtx context.Context

	// http
	http *http.Server

	// lock is used to safely access the client
	lock sync.RWMutex
}

// Server represents a listener
type Server listener

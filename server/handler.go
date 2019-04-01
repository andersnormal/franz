package server

import (
	"net"
	"net/http"
	"net/url"
)

// handler returns the redirector function for the HTTP server
func (s *Server) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

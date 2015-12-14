// Package server provides an extensible DDP server
package server

import (
	log "gopkg.in/inconshreveable/log15.v2"

	"golang.org/x/net/websocket"

	"net/http"
)

// Server is the DDP server
type Server struct {
	Log log.Logger

	listen string
	url    string

	server *websocket.Server
}

// NewServer creates a new server listening on the given address:port pair
func NewServer(listen string) *Server {
	s := &Server{
		Log:    log.New("pkg", "ddpd/server"),
		listen: listen,
		url:    "/websocket",
	}

	return s
}

// Run runs the server
func (s *Server) Run() error {
	s.Log.Debug("Running DDP server", "listen", s.listen)
	http.Handle(s.url, websocket.Handler(s.handler))
	err := http.ListenAndServe(s.listen, nil)
	return err
}

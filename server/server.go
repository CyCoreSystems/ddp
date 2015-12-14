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
		Log:    log.New("pkg", "ddp/server"),
		listen: listen,
		url:    "/websocket",
	}

	return s
}

// Run runs the server
func (s *Server) Run() error {
	s.Log.Debug("Running DDP server", "listen", s.listen)
	s.server = &websocket.Server{Handler: s.handler, Handshake: nil}
	http.Handle(s.url, s.server)
	err := http.ListenAndServe(s.listen, nil)
	return err
}

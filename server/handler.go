package server

import (
	"io"

	"golang.org/x/net/websocket"
)

func (s *Server) handler(ws *websocket.Conn) {

	s.Log.Debug("Got connection")

	io.Copy(ws, ws)
}

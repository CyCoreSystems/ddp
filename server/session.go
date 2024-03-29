package server

import (
	log "gopkg.in/inconshreveable/log15.v2"

	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
)

// Session represents a client connection
type Session struct {
	Log log.Logger

	ID string

	methodHandlerCtx context.Context
	conn             *websocket.Conn
}

package server

import (
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"

	"github.com/CyCoreSystems/ddp"
)

func (s *Session) ping() {

	s.Log.Debug("Sending ping", "id", "12")

	p := &ddp.Message{
		Type: "ping",
		ID:   "12", //TODO: use UUID
	}

	websocket.JSON.Send(s.conn, p)
}

func (s *Session) pong() {

	s.Log.Debug("Sending pong", "id", "12")

	p := &ddp.Message{
		Type: "pong",
		ID:   "12", //TODO: use UUID
	}

	websocket.JSON.Send(s.conn, p)
}

func (s *Session) pingHandler(ctx context.Context, msg *ddp.Message) error {
	s.Log.Debug("Got Ping", "id", msg.ID)
	s.pong()
	return nil
}

func (s *Session) pongHandler(ctx context.Context, msg *ddp.Message) error {
	s.Log.Debug("Got Pong", "id", msg.ID)
	return nil
}

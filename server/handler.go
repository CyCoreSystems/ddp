package server

import (
	"io"

	"golang.org/x/net/context"
	"golang.org/x/net/websocket"

	"github.com/CyCoreSystems/ddp"
)

func (s *Server) handler(ws *websocket.Conn) {
	s.Log.Debug("Got websocket connection")

	var connect ddp.Connect

	websocket.JSON.Receive(ws, &connect)

	s.Log.Debug("Got connect request", "connect", connect)

	connected := &ddp.Connected{
		Session: "1", //TODO: uuid generate session ID
	}

	websocket.JSON.Send(ws, connected)

	session := &Session{
		Log:  s.Log.New("session", "1"), //TODO: use previous ID
		ID:   "1",                       //TODO: use previous ID
		conn: ws,
	}

	session.router()
}

func (s *Session) router() {

	ctx := context.Background()
	ctx = ddp.RegisterHandler(ctx, "ping", s.pingHandler)
	ctx = ddp.RegisterHandler(ctx, "pong", s.pongHandler)

	s.ping()

	for {
		var msg ddp.Message
		if err := websocket.JSON.Receive(s.conn, &msg); err != nil {
			if err == io.EOF { // End of Connection?
				continue
			}
			s.Log.Error("Error reading JSON payload", "error", err)
			continue
		}

		if msg.Type == "" {
			continue
		}

		if err := ddp.CallHandler(ctx, msg.Type, &msg); err != nil {

			//TODO: respond with error

			s.Log.Error("Error calling handler", "error", err, "handler", msg.Type)
		}
	}

}

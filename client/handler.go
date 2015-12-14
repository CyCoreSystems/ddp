package client

import (
	"io"

	"golang.org/x/net/context"
	"golang.org/x/net/websocket"

	"github.com/CyCoreSystems/ddp"
)

func (s *Client) router() {

	ctx := context.Background()
	ctx = ddp.RegisterHandler(ctx, "ping", s.pingHandler)
	ctx = ddp.RegisterHandler(ctx, "pong", s.pongHandler)

	for {
		var msg ddp.Message
		if err := websocket.JSON.Receive(s.conn, &msg); err != nil {

			if err == io.EOF {
				continue
			}

			s.Log.Error("Error receiving JSON body", "error", err)
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

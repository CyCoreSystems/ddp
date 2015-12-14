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
		msg := make(ddp.Message)
		if err := websocket.JSON.Receive(s.conn, &msg); err != nil {

			if err == io.EOF {
				continue
			}

			s.Log.Error("Error receiving JSON body", "error", err)
			continue
		}

		t := msg.Type()
		if t == "" {
			continue
		}

		if err := ddp.CallHandler(ctx, t, msg); err != nil {

			//TODO: respond with error

			s.Log.Error("Error calling handler", "error", err, "handler", msg.Type)
		}
	}

}

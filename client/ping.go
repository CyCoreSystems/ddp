package client

import (
	"golang.org/x/net/context"
	"golang.org/x/net/websocket"

	"github.com/CyCoreSystems/ddp"
)

func (c *Client) ping() {
	c.Log.Debug("Sending ping", "id", "12")
	p := ddp.NewMessage("ping", "12") //TODO: use UUID
	websocket.JSON.Send(c.conn, p)
}

func (c *Client) pong(pingID string) {
	c.Log.Debug("Sending pong", "id", pingID)
	p := ddp.NewMessage("pong", pingID)
	websocket.JSON.Send(c.conn, p)
}

func (c *Client) pingHandler(ctx context.Context, msg ddp.Message) error {
	c.Log.Debug("Got Ping", "id", msg.ID())
	c.pong(msg.ID())
	return nil
}

func (c *Client) pongHandler(ctx context.Context, msg ddp.Message) error {
	c.Log.Debug("Got Pong", "id", msg.ID())
	return nil
}

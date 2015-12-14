package client

import (
	"golang.org/x/net/websocket"

	"github.com/CyCoreSystems/ddp"
)

func (c *Client) connect() {

	connect := &ddp.Connect{
		Version: "1",
	}
	c.Log.Debug("Sending connect message", "connect", connect)

	websocket.JSON.Send(c.conn, connect)

	var data map[string]interface{} = make(map[string]interface{})

	websocket.JSON.Receive(c.conn, &data)

	c.Log.Debug("Received connect response", "response", data)

	if sessionID, ok := data["session"].(string); ok {
		c.Log.Debug("Got session for client", "session", sessionID)
		c.SessionID = sessionID
	} else if version, ok := data["version"]; ok {
		c.Log.Debug("Got failed connect response", "version", version)
	}
}

package client

import (
	log "gopkg.in/inconshreveable/log15.v2"

	"golang.org/x/net/context"
	"golang.org/x/net/websocket"
)

// Client represents the DDP client
type Client struct {
	Log log.Logger

	url string

	conn *websocket.Conn

	ctx     context.Context
	closeFn func()
}

// NewClient creates a new DDP client
func NewClient(url string) *Client {
	cl := &Client{
		url: url,
		Log: log.New("pkg", "ddpd/client"),
	}

	cl.ctx, cl.closeFn = context.WithCancel(context.Background())

	return cl
}

// Run runs the DDP client
func (c *Client) Run() error {

	c.Log.Debug("Connecting to DDP server", "url", c.url)

	conn, err := websocket.Dial(c.url, "", "http://localhost/")
	if err != nil {
		return err
	}
	c.conn = conn

	c.Log.Debug("Waiting for DDP client to stop")
	<-c.ctx.Done()

	return nil
}

// Close stops the DDP client
func (c *Client) Close() {
	c.Log.Debug("DDP Client got close")
	c.closeFn()
}

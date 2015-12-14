package server

import (
	"github.com/CyCoreSystems/ddp"

	"golang.org/x/net/context"
)

func (s *Session) helloRpcHandler(ctx context.Context, msg ddp.Message) error {

	s.Log.Debug("Got Hello RPC", "msg", msg)

	return nil
}

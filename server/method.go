package server

import (
	"golang.org/x/net/context"

	"github.com/CyCoreSystems/ddp"
)

// the method handler implements a form of RPC

func (s *Session) methodHandler(ctx context.Context, msg ddp.Message) error {
	return nil
}

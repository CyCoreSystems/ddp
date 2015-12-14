package server

import (
	"golang.org/x/net/context"

	"github.com/CyCoreSystems/ddp"
)

// RegisterMethod registers an RPC method
func (s *Session) RegisterMethod(name string, handler ddp.Handler) {
	s.methodHandlerCtx = ddp.RegisterHandler(s.methodHandlerCtx, "rpc."+name, handler)
}

// the method handler implements a form of RPC

func (s *Session) methodHandler(ctx context.Context, msg ddp.Message) error {
	methodName := msg["method"].(string)
	return ddp.CallHandler(s.methodHandlerCtx, "rpc."+methodName, msg)
}

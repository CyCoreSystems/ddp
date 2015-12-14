package ddp

import (
	"golang.org/x/net/context"
)

//go:generate genrouter -type context -fntype Handler -keytype string .

type Handler func(ctx context.Context, msg *Message) error

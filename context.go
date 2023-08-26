package card

import "context"

type Context interface {
	context.Context
	Server
}

type ServerContext interface {
	context.Context
	Plugin
}

package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var OrderPS = wire.NewSet(NewOrderHTTPServer, NewOrderGRPCServer)

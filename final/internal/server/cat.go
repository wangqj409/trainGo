package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var CatPS = wire.NewSet(NewCatHTTPServer, NewCatGRPCServer)

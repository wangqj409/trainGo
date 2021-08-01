package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var GoodsPS = wire.NewSet(NewGoodsHTTPServer, NewGoodsGRPCServer)

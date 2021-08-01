package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var SupplyPS = wire.NewSet(NewSupplyHTTPServer, NewSupplyGRPCServer)
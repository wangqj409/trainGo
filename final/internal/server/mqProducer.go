package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var MqPS = wire.NewSet(NewMqProducerGRPCServer)

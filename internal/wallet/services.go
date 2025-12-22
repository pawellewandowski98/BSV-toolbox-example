package wallet

import (
	"log/slog"

	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/defs"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/services"
)

func CreateServices(network defs.BSVNetwork, log *slog.Logger) *services.WalletServices {
	servicesCfg := defs.DefaultServicesConfig(network)

	log.Info("Initializing services...", "network", network)

	s := services.New(log, servicesCfg)

	return s
}

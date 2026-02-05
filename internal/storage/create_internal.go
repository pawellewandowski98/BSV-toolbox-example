package storage

import (
	"SimpleScripts/internal/wallet"
	"context"
	"fmt"
	"log/slog"

	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/defs"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/services"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/storage"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/wdk"
)

func CreateInternal(network defs.BSVNetwork, storageKey string, onTxBroadcasted chan defs.TransactionStatusUpdate, log *slog.Logger) (*services.WalletServices, *storage.Provider, error) {
	s := wallet.CreateServices(network, log)

	log.Info("Creating storage provider...", "network", network)

	provider, err := storage.NewGORMProvider(network, s, storage.WithFeeModel(defs.FeeModel{
		Type:  defs.SatPerKB,
		Value: 100,
	}), storage.WithBackgroundBroadcasterChannel(onTxBroadcasted))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create storage provider: %w", err)
	}

	log.Info("Storage provider created successfully", "provider", provider)

	serverPrivateKey, err := wdk.IdentityKey(storageKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create server identity key: %w", err)
	}

	_, err = provider.Migrate(context.Background(), "go-storage-server", serverPrivateKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to migrate storage: %w", err)
	}

	return s, provider, nil
}

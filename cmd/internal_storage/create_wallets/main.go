package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/storage"
	"context"

	"SimpleScripts/internal/wallet"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW: create wallets to internal storage")

	_, provider, err := storage.CreateInternal(_const.Network, _const.ServerPrivateKeyHex, &log.Logger)

	aliceWallet, bobWallet, err := wallet.CreateAliceAndBobWallets(provider, _const.Network)
	if err != nil {
		log.Error("Failed to create Alice and Bob wallets", "error", err)
		return
	}

	log.Info("Wallets created successfully")
	response, err := aliceWallet.Wallet.GetVersion(context.Background(), nil, "alice")
	if err != nil {
		log.Error("Failed to get version", "error", err)
		return
	}

	log.Info("Alice wallet version retrieved", "version", response.Version)

	response, err = bobWallet.Wallet.GetVersion(context.Background(), nil, "bob")
	if err != nil {
		log.Error("Failed to get version", "error", err)
		return
	}

	log.Info("Bob wallet version retrieved", "version", response.Version)
	log.Complete("WORKFLOW COMPLETED: all wallets are operational")
}

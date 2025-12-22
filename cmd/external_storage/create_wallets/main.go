package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/wallet"
	"context"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW: create wallets to external storage")

	serverWallet, aliceWallet, bobWallet, err := wallet.CreateWalletsToExternalStorage(_const.Network)
	if err != nil {
		log.Error("Failed to create wallets", "error", err)
		return
	}

	defer func() {
		serverWallet.Cleanup()
		aliceWallet.Cleanup()
		bobWallet.Cleanup()
	}()

	log.Info("Wallets created successfully")
	response, err := serverWallet.Wallet.GetVersion(context.Background(), nil, "server")
	if err != nil {
		log.Error("Failed to get version", "error", err)
		return
	}

	log.Info("Server wallet version retrieved", "version", response.Version)

	response, err = aliceWallet.Wallet.GetVersion(context.Background(), nil, "alice")
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
	log.Complete("WORKFLOW COMPLETED: all wallets are operational ✅  ✅  ✅")
}

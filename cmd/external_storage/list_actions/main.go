package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/wallet"
	"context"

	sdkWallet "github.com/bsv-blockchain/go-sdk/wallet"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW: list outputs from wallets in external storage")

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

	reference := "random_reference_1770285883037501000"

	response, err := aliceWallet.Wallet.ListActions(context.Background(), sdkWallet.ListActionsArgs{
		Reference: &reference,
	}, "server")
	if err != nil {
		log.Error("Failed to list actions", "error", err)
		return
	}

	log.Info("Alice wallet actions retrieved", "number of actions", len(response.Actions))
	log.Info("Alice wallet actions retrieved", "actions", response.Actions)

	for i, action := range response.Actions {
		log.Info("Action retrieved", "action number", i, "action", action)
	}

	log.Complete("WORKFLOW COMPLETED: actions listed successfully")
}

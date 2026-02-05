package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/wallet"
	"context"

	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/wdk"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW: list transactions from wallets in external storage")

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

	response, err := aliceWallet.Wallet.ListTransactions(context.Background(), wdk.ListTransactionsArgs{
		//Reference: "random_reference_1770122402399307000",
	}, "server")
	if err != nil {
		log.Error("Failed to list outputs", "error", err)
		return
	}

	log.Info("Alice wallet outputs retrieved", "outputs", response.Transactions)

	log.Complete("WORKFLOW COMPLETED: actions listed successfully")
}

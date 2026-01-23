package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/tx"

	"SimpleScripts/internal/wallet"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW: send TX in internal storage")

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

	// NOTE: You need to have some outputs ready to spend
	//s := wallet.CreateServices(_const.Network, &log.Logger)
	//tx.Internalize(s, aliceWallet.Wallet, _const.TxIDToInternalize, &log.Logger)

	err = tx.SendTxWithExternalStorage(aliceWallet, bobWallet, 160, &log.Logger)
	if err != nil {
		log.Error("Failed to send TX from Alice to Bob", "error", err)
		return
	}

	log.Complete("WORKFLOW COMPLETED: all wallets are operational")
}

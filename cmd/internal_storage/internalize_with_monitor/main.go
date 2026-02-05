package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/monitor"
	"SimpleScripts/internal/storage"
	"SimpleScripts/internal/tx"
	"context"
	"os"
	"os/signal"
	"syscall"

	"SimpleScripts/internal/wallet"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW: internalize TX in internal storage")

	onTxBroadcasted, onTxProven := monitor.PrepareChannels(100)

	s, provider, err := storage.CreateInternal(_const.Network, _const.ServerPrivateKeyHex, onTxBroadcasted, &log.Logger)

	aliceWallet, bobWallet, err := wallet.CreateAliceAndBobWallets(provider, _const.Network)
	if err != nil {
		log.Error("Failed to create Alice and Bob wallets", "error", err)
		return
	}

	_, cleanup, err := monitor.RunMonitor(context.Background(), provider, onTxBroadcasted, onTxProven, &log.Logger)

	defer func() {
		cleanup()
	}()

	log.Info("Wallets created successfully")

	// NOTE: Uncomment to generate BRC-29 address for transfer, send funds to it, and then internalize the TX
	//address, err := tx.GenerateAddressForTransfer(aliceWallet.PrivKey)
	//if err != nil {
	//	log.Error("Failed to generate BRC-29 address", "error", err)
	//	return
	//}
	//
	//log.Info("Generated BRC-29 address for transfer", "address", address.AddressString)

	for _, id := range _const.AliceTxIDsToInternalize {
		tx.Internalize(s, aliceWallet.Wallet, id, &log.Logger)
	}
	for _, id := range _const.BobTxIDsToInternalize {
		tx.Internalize(s, bobWallet.Wallet, id, &log.Logger)
	}

	log.Complete("WORKFLOW COMPLETED: all wallets are operational")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}

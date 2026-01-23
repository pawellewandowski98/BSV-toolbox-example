package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/tx"

	"SimpleScripts/internal/wallet"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW: internalize TX in internal storage")

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

	s := wallet.CreateServices(_const.Network, &log.Logger)

	// NOTE: Uncomment to generate BRC-29 address for transfer, send funds to it, and then internalize the TX
	address, err := tx.GenerateAddressForTransfer(aliceWallet.PrivKey)
	if err != nil {
		log.Error("Failed to generate BRC-29 address", "error", err)
		return
	}

	log.Info("Alice address for transfers", "address", address.AddressString)

	address, err = tx.GenerateAddressForTransfer(bobWallet.PrivKey)
	if err != nil {
		log.Error("Failed to generate BRC-29 address", "error", err)
		return
	}

	log.Info("Bob address for transfers", "address", address.AddressString)

	for _, id := range _const.AliceTxIDsToInternalize {
		tx.Internalize(s, aliceWallet, id, &log.Logger)
	}
	for _, id := range _const.BobTxIDsToInternalize {
		tx.Internalize(s, bobWallet, id, &log.Logger)
	}

	//tx.Internalize(s, bobWallet, "01a2f814bd6a09abc150129f81a9ae548c0a978e69e7e6d6c32ce34dcf303213", &log.Logger)

	log.Complete("WORKFLOW COMPLETED: all wallets are operational")
}

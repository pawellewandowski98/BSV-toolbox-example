package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/wallet"
	"context"
	"fmt"

	"github.com/bsv-blockchain/go-sdk/transaction"
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

	response, err := serverWallet.Wallet.ListOutputs(context.Background(), sdkWallet.ListOutputsArgs{
		Include: sdkWallet.OutputIncludeEntireTransactions,
	}, "server")
	if err != nil {
		log.Error("Failed to list outputs", "error", err)
		return
	}

	log.Info("Server wallet outputs retrieved", "outputs", response)

	response, err = aliceWallet.Wallet.ListOutputs(context.Background(), sdkWallet.ListOutputsArgs{
		Include: sdkWallet.OutputIncludeEntireTransactions,
	}, "alice")
	if err != nil {
		log.Error("Failed to list outputs", "error", err)
		return
	}

	log.Info("Alice wallet outputs retrieved", "outputs", response)

	b, err := transaction.NewBeefFromBytes(response.BEEF)
	if err != nil {
		log.Error("Failed to parse beef from Alice's outputs", "error", err)
		return
	}

	fmt.Println("Beef from alice:", b)

	txs := b.GetValidTxids()
	fmt.Println("Alice's outputs include transactions with the following TXIDs: ", txs)

	response, err = bobWallet.Wallet.ListOutputs(context.Background(), sdkWallet.ListOutputsArgs{
		Include: sdkWallet.OutputIncludeEntireTransactions,
	}, "bob")
	if err != nil {
		log.Error("Failed to list outputs", "error", err)
		return
	}

	log.Info("Bob wallet outputs retrieved", "outputs", response)

	b, err = transaction.NewBeefFromBytes(response.BEEF)
	if err != nil {
		log.Error("Failed to parse beef from bob's outputs", "error", err)
		return
	}

	fmt.Println("Beef from bob:", b)

	txs = b.GetValidTxids()
	fmt.Println("Bob's outputs include transactions with the following TXIDs: ", txs)

	log.Complete("WORKFLOW COMPLETED: outputs listed successfully")
}

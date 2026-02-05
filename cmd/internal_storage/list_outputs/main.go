package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/storage"
	"SimpleScripts/internal/wallet"
	"context"
	"fmt"

	"github.com/bsv-blockchain/go-sdk/transaction"
	sdkWallet "github.com/bsv-blockchain/go-sdk/wallet"
	"github.com/go-softwarelab/common/pkg/to"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW:  list outputs in internal storage")

	_, provider, err := storage.CreateInternal(_const.Network, _const.ServerPrivateKeyHex, nil, &log.Logger)

	aliceWallet, bobWallet, err := wallet.CreateAliceAndBobWallets(provider, _const.Network)
	if err != nil {
		log.Error("Failed to create Alice and Bob wallets", "error", err)
		return
	}

	log.Info("Wallets created successfully")

	response, err := aliceWallet.Wallet.ListOutputs(context.Background(), sdkWallet.ListOutputsArgs{
		Include: sdkWallet.OutputIncludeEntireTransactions,
		Limit:   to.Ptr(uint32(100)),
	}, "alice")
	if err != nil {
		log.Error("Failed to list outputs", "error", err)
		return
	}

	log.Info("Alice wallet outputs retrieved", "outputs", response.Outputs)

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

	log.Info("Bob wallet outputs retrieved", "outputs", response.Outputs)

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

package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/tx"
	"SimpleScripts/internal/wallet"
	"context"
	"fmt"

	"github.com/bsv-blockchain/go-sdk/transaction"
	sdkWallet "github.com/bsv-blockchain/go-sdk/wallet"
)

func main() {
	log := logger.New()

	log.Start("STARTING WORKFLOW: create tx from input from wallets in external storage")

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

	s := wallet.CreateServices(_const.Network, &log.Logger)

	response, err := aliceWallet.Wallet.ListOutputs(context.Background(), sdkWallet.ListOutputsArgs{
		Include: sdkWallet.OutputIncludeEntireTransactions,
	}, "alice")
	if err != nil {
		log.Error("Failed to list outputs", "error", err)
		return
	}

	log.Info("Alice wallet outputs retrieved", "outputs", response.Outputs)

	//b, err := transaction.NewBeefFromBytes(response.BEEF)
	//if err != nil {
	//	log.Error("Failed to parse beef from Alice's outputs", "error", err)
	//	return
	//}
	//
	//fmt.Println("Beef from alice:", b)
	//
	//txs := b.GetValidTxids()
	//fmt.Println("Alice's outputs include transactions with the following TXIDs: ", txs)

	amountToSend := uint64(10)

	//output := response.Outputs[0]
	//if output.Satoshis != amountToSend {
	//	for _, o := range response.Outputs {
	//		fmt.Println("Checking output with satoshis:", o.Satoshis)
	//		if o.Satoshis == amountToSend {
	//			output = o
	//			break
	//		}
	//	}
	//
	//	if output.Satoshis != amountToSend {
	//		log.Error("No suitable output with more than 100 satoshi found in Alice's wallet")
	//		return
	//	}
	//}

	// bob -> alice
	txID := "388d77343fa4940eaf9c79ce9c3b266cddd4e14f2392ac586a2fe6a28bd35a84"

	outpoint, err := transaction.OutpointFromString(txID + ".0")
	if err != nil {
		log.Error("failed to parse outpoint: %w", err)
		return
	}

	log.Info("Selected output for spending", "output", outpoint)

	beef, err := s.GetBEEF(context.Background(), txID, nil)
	if err != nil {
		panic(fmt.Errorf("failed to get BEEF for txID %s: %w", txID, err))
	}

	beefBytes, err := beef.Bytes()
	if err != nil {
		log.Error("Failed to get beef bytes", "error", err)
		return
	}

	// alice -> bob
	err = tx.SendTxWithInputWithExternalStorage(bobWallet, aliceWallet, beefBytes, *outpoint, amountToSend, &log.Logger)
	if err != nil {
		log.Error("Failed to send TX from Alice to Bob", "error", err)
	}

	log.Complete("WORKFLOW COMPLETED: outputs listed successfully")
}

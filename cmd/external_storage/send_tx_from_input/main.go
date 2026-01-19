package main

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/logger"
	"SimpleScripts/internal/tx"
	"SimpleScripts/internal/wallet"
	"context"
	"fmt"

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

	amountToSend := uint64(200)

	output := response.Outputs[0]
	if output.Satoshis < amountToSend {
		for _, o := range response.Outputs {
			fmt.Println("Checking output with satoshis:", o.Satoshis)
			if o.Satoshis > amountToSend {
				output = o
				break
			}
		}

		if output.Satoshis <= amountToSend {
			log.Error("No suitable output with more than 100 satoshi found in Alice's wallet")
			return
		}
	}

	log.Info("Selected output for spending", "output", output)

	err = tx.SendTxWithInputWithExternalStorage(aliceWallet, bobWallet, response.BEEF, output, amountToSend, &log.Logger)
	if err != nil {
		log.Error("Failed to send TX from Alice to Bob", "error", err)
	}

	log.Complete("WORKFLOW COMPLETED: outputs listed successfully")
}

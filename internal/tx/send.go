package tx

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/wallet"
	"context"
	"fmt"
	"log/slog"

	"github.com/bsv-blockchain/go-sdk/transaction/template/p2pkh"
	sdkWallet "github.com/bsv-blockchain/go-sdk/wallet"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/brc29"
	"github.com/go-softwarelab/common/pkg/to"
)

func SendTx(sender, recipient *wallet.WalletWithKeys, log *slog.Logger) error {
	address, err := brc29.AddressForCounterparty(
		sender.PrivKey,
		_const.KeyID,
		recipient.PubKey,
		brc29.WithMainNet(),
	)
	if err != nil {
		return fmt.Errorf("failed to generate BRC29 address: %w", err)
	}

	log.Info("Generated BRC29 address", "address", address.AddressString)

	lockingScript, err := p2pkh.Lock(address)
	if err != nil {
		return fmt.Errorf("failed to create locking script: %w", err)
	}

	log.Info("Created locking script", "lockingScript", lockingScript)

	createArgs := sdkWallet.CreateActionArgs{
		Description: "TX example",
		Outputs: []sdkWallet.CreateActionOutput{
			{
				LockingScript:     lockingScript.Bytes(),
				Satoshis:          uint64(1),
				OutputDescription: "Payment to BRC29 address",
				Tags:              []string{"payment", "example"},
			},
		},
		Labels: []string{"create_action_example"},
		Options: &sdkWallet.CreateActionOptions{
			AcceptDelayedBroadcast: to.Ptr(false),
		},
	}

	log.Info("Creating transaction to send 1 satoshi", "description", createArgs.Description)

	result, err := sender.Wallet.CreateAction(context.Background(), createArgs, "test_originator")
	if err != nil {
		return fmt.Errorf("failed to create action: %w", err)
	}

	log.Info("CreateAction successful", "result", *result)

	if result.Txid.String() != "" {
		log.Info("Transaction successfully created and broadcast", "txID", result.Txid.String())

		if len(result.SendWithResults) > 0 {
			log.Info("Broadcast status", result.SendWithResults[0].Status)
		}
	}

	return nil
}

func SendTxWithExternalStorage(sender, recipient *wallet.WalletForExternalStorageWithKeys, log *slog.Logger) error {
	senderWalletWithKeys := &wallet.WalletWithKeys{
		Wallet:  sender.Wallet,
		PrivKey: sender.PrivKey,
		PubKey:  sender.PubKey,
	}

	recipientWalletWithKeys := &wallet.WalletWithKeys{
		Wallet:  recipient.Wallet,
		PrivKey: recipient.PrivKey,
		PubKey:  recipient.PubKey,
	}

	err := SendTx(senderWalletWithKeys, recipientWalletWithKeys, log)
	if err != nil {
		return err
	}

	return nil
}

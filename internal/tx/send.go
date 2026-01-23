package tx

import (
	_const "SimpleScripts/const"
	"SimpleScripts/internal/wallet"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/bsv-blockchain/go-sdk/transaction"
	sdkWallet "github.com/bsv-blockchain/go-sdk/wallet"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/brc29"
	"github.com/go-softwarelab/common/pkg/to"
)

func SendTx(sender, recipient *wallet.WalletWithKeys, amount uint64, log *slog.Logger) error {
	//address, err := brc29.AddressForCounterparty(
	//	sender.PrivKey,
	//	_const.KeyID,
	//	recipient.PubKey,
	//	brc29.WithMainNet(),
	//)
	//if err != nil {
	//	return fmt.Errorf("failed to generate BRC29 address: %w", err)
	//}
	//
	//log.Info("Generated BRC29 address", "address", address.AddressString)
	//
	//lockingScript, err := p2pkh.Lock(address)
	//if err != nil {
	//	return fmt.Errorf("failed to create locking script: %w", err)
	//}

	lockingScript, err := brc29.LockForCounterparty(sender.PrivKey, _const.KeyID, recipient.PubKey)
	if err != nil {
		return fmt.Errorf("failed to create locking script: %w", err)
	}

	log.Info("Created locking script", "lockingScript", lockingScript)

	createArgs := sdkWallet.CreateActionArgs{
		Description: "TX example",
		Outputs: []sdkWallet.CreateActionOutput{
			{
				LockingScript:     lockingScript.Bytes(),
				Satoshis:          50,
				OutputDescription: "Payment to BRC29 address",
				Tags:              []string{"payment", "example"},
			},
		},
		Labels: []string{"create_action_example"},
		Options: &sdkWallet.CreateActionOptions{
			AcceptDelayedBroadcast: to.Ptr(true),
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

func SendTxWithExternalStorage(sender, recipient *wallet.WalletForExternalStorageWithKeys, amount uint64, log *slog.Logger) error {
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

	err := SendTx(senderWalletWithKeys, recipientWalletWithKeys, amount, log)
	if err != nil {
		return err
	}

	return nil
}
func SendTxWithInputWithExternalStorage(sender, recipient *wallet.WalletForExternalStorageWithKeys, inputBeef []byte, outpoint transaction.Outpoint, amount uint64, log *slog.Logger) error {
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

	err := SendTxWithInput(senderWalletWithKeys, recipientWalletWithKeys, inputBeef, outpoint, amount, log)
	if err != nil {
		return err
	}

	return nil
}

func SendTxWithInput(sender, recipient *wallet.WalletWithKeys, inputBeef []byte, outpoint transaction.Outpoint, amount uint64, log *slog.Logger) error {
	//address, err := brc29.AddressForCounterparty(
	//	sender.PrivKey,
	//	_const.KeyID,
	//	recipient.PubKey,
	//	brc29.WithMainNet(),
	//)
	//if err != nil {
	//	return fmt.Errorf("failed to generate BRC29 address: %w", err)
	//}
	//
	//log.Info("Generated BRC29 address", "address", address.AddressString)

	//lockingScript, err := p2pkh.Lock(address)
	//if err != nil {
	//	return fmt.Errorf("failed to create locking script: %w", err)
	//}

	//lockingScript, err := brc29.LockForCounterparty(sender.PrivKey, _const.KeyID, recipient.PubKey)
	//if err != nil {
	//	return fmt.Errorf("failed to create locking script: %w", err)
	//}

	//log.Info("Created locking script", "lockingScript", lockingScript)

	// alice, bob
	//unlocker, err := brc29.Unlock(sender.PubKey, _const.KeyID, recipient.PrivKey)
	//if err != nil {
	//	return fmt.Errorf("failed to create unlocker: %w", err)
	//}

	// bob, alice
	unlocker, err := brc29.Unlock(sender.PubKey, _const.KeyID, recipient.PrivKey)
	if err != nil {
		return fmt.Errorf("failed to create unlocker: %w", err)
	}

	createArgs := sdkWallet.CreateActionArgs{
		InputBEEF: inputBeef,
		Inputs: []sdkWallet.CreateActionInput{
			{
				Outpoint:              outpoint,
				InputDescription:      "funds source input",
				UnlockingScriptLength: unlocker.EstimateLength(nil, 0),
			},
		},
		// old code
		Description: "TX example",
		//Outputs: []sdkWallet.CreateActionOutput{
		//	{
		//		LockingScript:     lockingScript.Bytes(),
		//		Satoshis:          amount,
		//		OutputDescription: "Payment to BRC29 address",
		//		Tags:              []string{"payment", "example"},
		//	},
		//},
		Labels: []string{"create_action_example"},
		Options: &sdkWallet.CreateActionOptions{
			AcceptDelayedBroadcast: to.Ptr(true),
		},
	}

	log.Info("Creating transaction to send 1 satoshi", "description", createArgs.Description)

	result, err := sender.Wallet.CreateAction(context.Background(), createArgs, "test_originator")
	if err != nil {
		return fmt.Errorf("failed to create action: %w", err)
	}

	log.Info("CreateAction successful", "result", *result)

	if result.SignableTransaction == nil {
		return errors.New("signable transaction not found after create action")
	}

	if result.Txid.String() != "" {
		log.Info("Transaction successfully created", "txID", result.Txid.String())

		if len(result.SendWithResults) > 0 {
			log.Info("Broadcast status", result.SendWithResults[0].Status)
		}
	}

	txBeef, txHash, err := transaction.NewBeefFromAtomicBytes(result.SignableTransaction.Tx)
	if err != nil {
		return fmt.Errorf("error parsing signable transaction: %w", err)
	}
	tx := txBeef.FindAtomicTransactionByHash(txHash)

	unlockingScript, err := unlocker.Sign(tx, 0)
	if err != nil {
		return fmt.Errorf("error unlocking funding input: %w", err)
	}

	signArgs := sdkWallet.SignActionArgs{
		Reference: result.SignableTransaction.Reference,
		Spends: map[uint32]sdkWallet.SignActionSpend{
			0: {
				UnlockingScript: unlockingScript.Bytes(),
			},
		},
		Options: &sdkWallet.SignActionOptions{
			AcceptDelayedBroadcast: to.Ptr(true),
		},
	}
	sar, err := sender.Wallet.SignAction(context.Background(), signArgs, "test_originator")
	if err != nil {
		return fmt.Errorf("error signing and broadcasting transaction: %w", err)
	}

	log.Info("successful transaction", "txid", sar.Txid.String())

	return nil
}

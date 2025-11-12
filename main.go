package main

import (
	"SimpleScripts/utils"
	"context"
	"fmt"
	"log/slog"

	"github.com/bsv-blockchain/go-sdk/chainhash"
	"github.com/bsv-blockchain/go-sdk/transaction/template/p2pkh"
	sdk "github.com/bsv-blockchain/go-sdk/wallet"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/brc29"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/defs"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/services"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/storage"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/wdk"
	"github.com/go-softwarelab/common/pkg/to"
)

var log *slog.Logger

func main() {
	log = slog.Default()
	network := defs.NetworkMainnet

	s, provider := setupServer(network)

	alice, _ := utils.CreateWallets(provider, network)

	// Step 1: Generate BRC29 address for Alice
	generateAddress(alice)

	// Step 2: Internalize transaction from faucet to Alice's wallet
	//internalizeTx(s, alice, "299ac36833a5ffe6ae30e4d9fcebd6328a5fd4e6cae5dc4d18bda95adc1bbad1")

	// Step 3: Send transaction from Alice to Bob
	//sendTx(alice, bob)
}

func setupServer(network defs.BSVNetwork) (s *services.WalletServices, provider *storage.Provider) {
	log = slog.Default()
	servicesCfg := defs.DefaultServicesConfig(network)

	log.Info("Initializing services...", "network", network)

	s = services.New(log, servicesCfg)

	log.Info("Starting services...", "network", network)
	log.Info("Creating storage provider...", "network", network)

	var err error
	provider, err = storage.NewGORMProvider(network, s)
	if err != nil {
		log.Error("Failed to create storage provider", "error", err)
		return
	}

	log.Info("Storage provider created successfully", "provider", provider)

	serverPrivateKey, err := wdk.IdentityKey("37dd01d7b2f0978c3f72b64289accac122ddff3cc9d59e9708522cb804d4ec08")
	if err != nil {
		log.Error("Failed to create server identity key", "error", err)
		return
	}

	_, err = provider.Migrate(context.Background(), "go-storage-server", serverPrivateKey)
	if err != nil {
		log.Error("Failed to migrate storage", "error", err)
		return
	}

	return
}

func generateAddress(side utils.WalletWithKeys) {
	keyID := brc29.KeyID{
		DerivationPrefix: utils.DefaultBase64Prefix,
		DerivationSuffix: utils.DefaultBase64Suffix,
	}

	_, aPubKey := sdk.AnyoneKey()

	address, err := brc29.AddressForSelf(
		aPubKey,
		keyID,
		side.PrivKey,
		brc29.WithMainNet(),
	)
	if err != nil {
		log.Error("Failed to generate BRC29 address", "error", err)
		return
	}

	log.Info("Generated BRC29 address", "address", address.AddressString)
}

func internalizeTx(s *services.WalletServices, side utils.WalletWithKeys, txID string) {
	txIDHash, err := chainhash.NewHashFromHex(txID)
	if err != nil {
		panic(fmt.Errorf("invalid txID: %w", err))
	}

	beef, err := s.GetBEEF(context.Background(), txID, nil)
	if err != nil {
		panic(fmt.Errorf("failed to get BEEF for txID %s: %w", txID, err))
	}

	log.Info("Fetched BEEF successfully", "beef", beef)

	atomicBeef, err := beef.AtomicBytes(txIDHash)
	if err != nil {
		panic(fmt.Errorf("failed to get atomic bytes for txID %s: %w", txID, err))
	}

	log.Info("Obtained atomic bytes successfully", "atomicBeef", atomicBeef)

	err = utils.InternalizeFromFaucet(context.Background(), atomicBeef, side.Wallet)
	if err != nil {
		panic(fmt.Errorf("failed to internalize tx: %w", err))
	}
}

func sendTx(sender, recipient utils.WalletWithKeys) {
	keyID := brc29.KeyID{
		DerivationPrefix: utils.DefaultBase64Prefix,
		DerivationSuffix: utils.DefaultBase64Suffix,
	}

	address, err := brc29.AddressForCounterparty(
		sender.PrivKey,
		keyID,
		recipient.PubKey,
		brc29.WithMainNet(),
	)
	if err != nil {
		log.Error("Failed to generate BRC29 address", "error", err)
		return
	}

	log.Info("Generated BRC29 address", "address", address.AddressString)

	lockingScript, err := p2pkh.Lock(address)
	if err != nil {
		log.Error("Failed to create locking script", "error", err)
		return
	}

	log.Info("Created locking script", "lockingScript", lockingScript)

	createArgs := sdk.CreateActionArgs{
		Description: "TX example",
		Outputs: []sdk.CreateActionOutput{
			{
				LockingScript:     lockingScript.Bytes(),
				Satoshis:          uint64(1),
				OutputDescription: "Payment to BRC29 address",
				Tags:              []string{"payment", "example"},
			},
		},
		Labels: []string{"create_action_example"},
		Options: &sdk.CreateActionOptions{
			AcceptDelayedBroadcast: to.Ptr(false),
		},
	}

	log.Info("Creating transaction to send 1 satoshi", "description", createArgs.Description)

	result, err := sender.Wallet.CreateAction(context.Background(), createArgs, "test_originator")
	if err != nil {
		log.Error("Failed to create action", "error", err)
		return
	}

	log.Info("CreateAction successful", "result", *result)

	if result.Txid.String() != "" {
		log.Info("Transaction successfully created and broadcast", "txID", result.Txid.String())

		if len(result.SendWithResults) > 0 {
			log.Info("Broadcast status", result.SendWithResults[0].Status)
		}
	}

}

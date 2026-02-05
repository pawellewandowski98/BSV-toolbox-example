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

	log.Start("STARTING WORKFLOW: create TX in internal storage")

	onTxBroadcasted, onTxProven := monitor.PrepareChannels(100)

	_, provider, err := storage.CreateInternal(_const.Network, _const.ServerPrivateKeyHex, onTxBroadcasted, &log.Logger)

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

	err = tx.SendTx(aliceWallet, bobWallet, 1, &log.Logger)
	if err != nil {
		log.Error("Failed to send TX from Alice to Bob", "error", err)
		return
	}

	log.Complete("WORKFLOW COMPLETED: all wallets are operational")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}

//
//func setupServer(network defs.BSVNetwork) (s *services.WalletServices, provider *storage.Provider) {
//	log = slog.Default()
//	servicesCfg := defs.DefaultServicesConfig(network)
//
//	log.Info("Initializing services...", "network", network)
//
//	s = services.New(log, servicesCfg)
//
//	log.Info("Starting services...", "network", network)
//	log.Info("Creating storage provider...", "network", network)
//
//	var err error
//	provider, err = storage.NewGORMProvider(network, s)
//	if err != nil {
//		log.Error("Failed to create storage provider", "error", err)
//		return
//	}
//
//	log.Info("Storage provider created successfully", "provider", provider)
//
//	serverPrivateKey, err := wdk.IdentityKey("37dd01d7b2f0978c3f72b64289accac122ddff3cc9d59e9708522cb804d4ec08")
//	if err != nil {
//		log.Error("Failed to create server identity key", "error", err)
//		return
//	}
//
//	_, err = provider.Migrate(context.Background(), "go-storage-server", serverPrivateKey)
//	if err != nil {
//		log.Error("Failed to migrate storage", "error", err)
//		return
//	}
//
//	return
//}
//
//func internalizeTx(s *services.WalletServices, side *wallet.WalletWithKeys, txID string) {
//
//	keyID := brc29.KeyID{
//		DerivationPrefix: utils.DefaultBase64Prefix,
//		DerivationSuffix: utils.DefaultBase64Suffix,
//	}
//
//	_, aPubKey := sdkWallet.AnyoneKey()
//
//	address, err := brc29.AddressForSelf(
//		aPubKey,
//		keyID,
//		side.PrivKey,
//		brc29.WithMainNet(),
//	)
//	if err != nil {
//		log.Error("Failed to generate BRC29 address", "error", err)
//		return
//	}
//
//	log.Info("Generated BRC29 address", "address", address.AddressString)
//
//	txIDHash, err := chainhash.NewHashFromHex(txID)
//	if err != nil {
//		panic(fmt.Errorf("invalid txID: %w", err))
//	}
//
//	beef, err := s.GetBEEF(context.Background(), txID, nil)
//	if err != nil {
//		panic(fmt.Errorf("failed to get BEEF for txID %s: %w", txID, err))
//	}
//
//	log.Info("Fetched BEEF successfully", "beef", beef)
//
//	atomicBeef, err := beef.AtomicBytes(txIDHash)
//	if err != nil {
//		panic(fmt.Errorf("failed to get atomic bytes for txID %s: %w", txID, err))
//	}
//
//	log.Info("Obtained atomic bytes successfully", "atomicBeef", atomicBeef)
//
//	err = utils.InternalizeFromFaucet(context.Background(), atomicBeef, side.Wallet, nil)
//	if err != nil {
//		panic(fmt.Errorf("failed to internalize tx: %w", err))
//	}
//}
//
//func sendTx(sender, recipent *wallet.WalletWithKeys) {
//	keyID := brc29.KeyID{
//		DerivationPrefix: utils.DefaultBase64Prefix,
//		DerivationSuffix: utils.DefaultBase64Suffix,
//	}
//
//	address, err := brc29.AddressForCounterparty(
//		sender.PrivKey,
//		keyID,
//		recipent.PubKey,
//		brc29.WithMainNet(),
//	)
//	if err != nil {
//		log.Error("Failed to generate BRC29 address", "error", err)
//		return
//	}
//
//	log.Info("Generated BRC29 address", "address", address.AddressString)
//
//	lockingScript, err := p2pkh.Lock(address)
//	if err != nil {
//		log.Error("Failed to create locking script", "error", err)
//		return
//	}
//
//	log.Info("Created locking script", "lockingScript", lockingScript)
//
//	createArgs := sdkWallet.CreateActionArgs{
//		Description: "TX example",
//		Outputs: []sdkWallet.CreateActionOutput{
//			{
//				LockingScript:     lockingScript.Bytes(),
//				Satoshis:          uint64(1),
//				OutputDescription: "Payment to BRC29 address",
//				Tags:              []string{"payment", "example"},
//			},
//		},
//		Labels: []string{"create_action_example"},
//		Options: &sdkWallet.CreateActionOptions{
//			AcceptDelayedBroadcast: to.Ptr(false),
//		},
//	}
//
//	log.Info("Creating transaction to send 1 satoshi", "description", createArgs.Description)
//
//	result, err := sender.Wallet.CreateAction(context.Background(), createArgs, "test_originator")
//	if err != nil {
//		log.Error("Failed to create action", "error", err)
//		return
//	}
//
//	log.Info("CreateAction successful", "result", *result)
//
//	if result.Txid.String() != "" {
//		log.Info("Transaction successfully created and broadcast", "txID", result.Txid.String())
//
//		if len(result.SendWithResults) > 0 {
//			log.Info("Broadcast status", result.SendWithResults[0].Status)
//		}
//	}
//
//}

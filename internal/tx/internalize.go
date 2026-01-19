package tx

import (
	_const "SimpleScripts/const"
	"context"
	"encoding/base64"
	"fmt"
	"log/slog"

	"github.com/bsv-blockchain/go-sdk/chainhash"
	sdk "github.com/bsv-blockchain/go-sdk/wallet"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/services"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/wallet"
)

func Internalize(s *services.WalletServices, side *wallet.Wallet, txID string, log *slog.Logger) {
	log.Info("Starting internalization process", "txID", txID)

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

	log.Info("Obtained atomic bytes successfully")

	err = InternalizeFromFaucet(context.Background(), atomicBeef, side)
	if err != nil {
		panic(fmt.Errorf("failed to internalize tx: %w", err))
	}
}

// DerivationBytesResult represents the result of derivation bytes calculation
type DerivationBytesResult struct {
	DerivationPrefix []byte `json:"derivationPrefix"`
	DerivationSuffix []byte `json:"derivationSuffix"`
}

// InternalizeFromFaucet is a helper function to internalize a transaction from the faucet
func InternalizeFromFaucet(ctx context.Context, atomicBeefBytes []byte, wallet sdk.Interface) error {
	paymentRemittance := DerivationParts()

	internalizeArgs := sdk.InternalizeActionArgs{
		Tx: atomicBeefBytes,
		Outputs: []sdk.InternalizeOutput{
			{
				OutputIndex:       0,
				Protocol:          "wallet payment",
				PaymentRemittance: paymentRemittance,
			},
		},
		Description: "internalize from faucet",
	}

	iar, err := wallet.InternalizeAction(ctx, internalizeArgs, "originator")
	if err != nil {
		return fmt.Errorf("failed to internalize action: %w", err)
	}

	fmt.Printf("InternalizeAction successful: %+v\n", *iar)

	return nil
}

// DerivationParts creates derivation parts with default prefix and suffix
func DerivationParts() *sdk.Payment {
	prefix := "" // empty string will use default base64 prefix
	suffix := "" // empty string will use default base64 suffix
	bytes := derivationBytes(prefix, suffix)

	_, publicKey := sdk.AnyoneKey()

	paymentRemittance := &sdk.Payment{
		DerivationPrefix:  bytes.DerivationPrefix,
		DerivationSuffix:  bytes.DerivationSuffix,
		SenderIdentityKey: publicKey,
	}

	return paymentRemittance
}

func derivationBytes(prefix string, suffix string) DerivationBytesResult {
	var derivationPrefix []byte
	var derivationSuffix []byte
	var err error

	if prefix == "" {
		prefix = _const.DefaultBase64Prefix
	}

	derivationPrefix, err = base64.StdEncoding.DecodeString(prefix)
	if err != nil {
		panic(fmt.Errorf("failed to decode default base64 prefix: %w", err))
	}

	if suffix == "" {
		suffix = _const.DefaultBase64Suffix
	}

	derivationSuffix, err = base64.StdEncoding.DecodeString(suffix)
	if err != nil {
		panic(fmt.Errorf("failed to decode default base64 suffix: %w", err))
	}

	fmt.Printf("Using Derivation Prefix: %x\n", derivationPrefix)
	fmt.Printf("Using Derivation Prefix: %x\n", derivationSuffix)

	return DerivationBytesResult{
		DerivationPrefix: derivationPrefix,
		DerivationSuffix: derivationSuffix,
	}
}

package utils

import (
	"context"
	"encoding/base64"
	"fmt"

	sdk "github.com/bsv-blockchain/go-sdk/wallet"
)

const (
	DefaultBase64Prefix = "SfKxPIJNgdI="
	DefaultBase64Suffix = "NaGLC6fMH50="
)

type PaymentRemittance struct {
	DerivationPrefix  []byte `json:"derivationPrefix"`
	DerivationSuffix  []byte `json:"derivationSuffix"`
	SenderIdentityKey string `json:"senderIdentityKey"`
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
		prefix = DefaultBase64Prefix
	}

	derivationPrefix, err = base64.StdEncoding.DecodeString(prefix)
	if err != nil {
		panic(fmt.Errorf("failed to decode default base64 prefix: %w", err))
	}

	if suffix == "" {
		suffix = DefaultBase64Suffix
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

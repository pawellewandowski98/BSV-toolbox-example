package tx

import (
	_const "SimpleScripts/const"
	"fmt"

	ec "github.com/bsv-blockchain/go-sdk/primitives/ec"
	"github.com/bsv-blockchain/go-sdk/script"
	"github.com/bsv-blockchain/go-sdk/wallet"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/brc29"
)

func GenerateAddressForTransfer(recipient *ec.PrivateKey) (*script.Address, error) {
	keyID := brc29.KeyID{
		DerivationPrefix: _const.DefaultBase64Prefix,
		DerivationSuffix: _const.DefaultBase64Suffix,
	}

	_, aPubKey := wallet.AnyoneKey()

	address, err := brc29.AddressForSelf(
		aPubKey,
		keyID,
		recipient,
		brc29.WithMainNet(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate BRC-29 address: %w", err)
	}

	return address, nil
}

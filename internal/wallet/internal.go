package wallet

import (
	_const "SimpleScripts/const"
	"fmt"

	ec "github.com/bsv-blockchain/go-sdk/primitives/ec"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/defs"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/wallet"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/wdk"
)

type WalletWithKeys struct {
	Wallet  *wallet.Wallet
	PrivKey *ec.PrivateKey
	PubKey  *ec.PublicKey
}

func CreateAliceAndBobWallets(s wdk.WalletStorageProvider, network defs.BSVNetwork) (alice, bob *WalletWithKeys, err error) {
	alice, err = CreateWalletWithStorageProvider(s, network, _const.AlicePrivateKey)
	if err != nil {
		return nil, nil, err
	}

	bob, err = CreateWalletWithStorageProvider(s, network, _const.BobPrivateKey)
	if err != nil {
		return nil, nil, err
	}

	return
}

func CreateWalletWithStorageProvider(s wdk.WalletStorageProvider, network defs.BSVNetwork, key string) (*WalletWithKeys, error) {
	priv, pub := ec.PrivateKeyFromBytes([]byte(key))
	if priv == nil {
		return nil, fmt.Errorf("cannot create private key from bytes")
	}

	w, err := wallet.New(network, priv, s)
	if err != nil {
		return nil, fmt.Errorf("cannot create wallet: %w", err)
	}

	//w, err := wallet.NewWithStorageFactory(network, priv, func() wdk.WalletStorageProvider { return s })
	//if err != nil {
	//	panic("Cannot create wallet")
	//}

	return &WalletWithKeys{
		Wallet:  w,
		PrivKey: priv,
		PubKey:  pub,
	}, nil
}

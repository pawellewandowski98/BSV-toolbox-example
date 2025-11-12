package utils

import (
	ec "github.com/bsv-blockchain/go-sdk/primitives/ec"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/defs"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/storage"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/wallet"
)

type WalletWithKeys struct {
	Wallet  *wallet.Wallet
	PrivKey *ec.PrivateKey
	PubKey  *ec.PublicKey
}

func CreateWallets(s *storage.Provider, network defs.BSVNetwork) (alice WalletWithKeys, bob WalletWithKeys) {
	alice = createWallet(s, network, "xprv9s21ZrQH143K3WYAquX13GWNfPShBx5XT98kBDMQpxz5p1EYJ8fsqwQCkKuJyB7")
	bob = createWallet(s, network, "xprv9s21ZrQH143K2hza6vc9EsRwv2sZTUNc9EHfZByqBaThBsrr3u8eKfFCt5KRUr721v")

	return
}

func createWallet(s *storage.Provider, network defs.BSVNetwork, key string) WalletWithKeys {
	priv, pub := ec.PrivateKeyFromBytes([]byte(key))
	if priv == nil {
		panic("Cannot create private key")
	}

	w, err := wallet.New(network, priv, s)
	if err != nil {
		panic("Cannot create wallet")
	}

	return WalletWithKeys{
		Wallet:  w,
		PrivKey: priv,
		PubKey:  pub,
	}
}

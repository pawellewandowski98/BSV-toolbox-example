package wallet

import (
	_const "SimpleScripts/const"
	"fmt"

	ec "github.com/bsv-blockchain/go-sdk/primitives/ec"
	sdkWallet "github.com/bsv-blockchain/go-sdk/wallet"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/defs"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/storage"
	"github.com/bsv-blockchain/go-wallet-toolbox/pkg/wallet"
)

type WalletForExternalStorageWithKeys struct {
	Wallet      *wallet.Wallet
	ProtoWallet *sdkWallet.CompletedProtoWallet
	PrivKey     *ec.PrivateKey
	PubKey      *ec.PublicKey

	Cleanup func()
}

func CreateWalletsToExternalStorage(network defs.BSVNetwork) (serverWallet, aliceWallet, bobWallet *WalletForExternalStorageWithKeys, err error) {
	serverKey, err := ec.PrivateKeyFromHex(_const.ServerPrivateKeyHex)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create private key: %w", err)
	}

	serverWallet, err = CreateWalletWithClient(network, serverKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create server wallet: %w", err)
	}

	aliceKey, _ := ec.PrivateKeyFromBytes([]byte(_const.AlicePrivateKey))
	if aliceKey == nil {
		return nil, nil, nil, fmt.Errorf("cannot create Alice private key")
	}

	aliceWallet, err = CreateWalletWithClient(network, aliceKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create Alice wallet: %w", err)
	}

	bobKey, _ := ec.PrivateKeyFromBytes([]byte(_const.BobPrivateKey))
	if bobKey == nil {
		return nil, nil, nil, fmt.Errorf("cannot create Bob private key")
	}

	bobWallet, err = CreateWalletWithClient(network, bobKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to create Bob wallet: %w", err)
	}

	return
}

func CreateWalletWithClient(network defs.BSVNetwork, key *ec.PrivateKey) (*WalletForExternalStorageWithKeys, error) {
	pw, err := sdkWallet.NewCompletedProtoWallet(key)
	wspc, cleanup, err := storage.NewClient(_const.StorageURL, pw)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage provider client: %w", err)
	}

	w, err := wallet.New(network, key, wspc)
	if err != nil {
		return nil, fmt.Errorf("failed to create server wallet: %w", err)
	}

	return &WalletForExternalStorageWithKeys{
		Wallet:      w,
		ProtoWallet: pw,
		PrivKey:     key,
		PubKey:      key.PubKey(),
		Cleanup:     cleanup,
	}, nil
}

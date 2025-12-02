package main

import "context"

type WalletRepository interface {
	GetWallet(ctx context.Context, id string) (*Wallet, error)
	SaveWallet(ctx context.Context, w *Wallet) error
}

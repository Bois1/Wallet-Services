package main

import (
	"context"
	"errors"
)

type FailingWalletRepo struct{}

func (f FailingWalletRepo) GetWallet(ctx context.Context, id string) (*Wallet, error) {
	return nil, errors.New("storage unavailable")
}

func (f FailingWalletRepo) SaveWallet(ctx context.Context, w *Wallet) error {
	return errors.New("storage unavailable")
}

package main

import (
	"context"
	"errors"
	"sync"
)

type InMemoryWalletRepo struct {
	wallets map[string]*Wallet
	mu      sync.RWMutex
}

func NewInMemoryWalletRepo() *InMemoryWalletRepo {
	return &InMemoryWalletRepo{
		wallets: make(map[string]*Wallet),
	}
}

func (r *InMemoryWalletRepo) GetWallet(ctx context.Context, id string) (*Wallet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if w, ok := r.wallets[id]; ok {

		copied := *w
		return &copied, nil
	}
	return nil, errors.New("wallet not found")
}

func (r *InMemoryWalletRepo) SaveWallet(ctx context.Context, w *Wallet) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	copied := *w
	r.wallets[w.ID] = &copied
	return nil
}

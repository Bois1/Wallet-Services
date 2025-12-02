package main

import (
	"context"
	"errors"
	"sync"
)

type TransferRequest struct {
	ID     string
	FromID string
	ToID   string
	Amount Money
}

type WalletService struct {
	repo       WalletRepository
	mu         sync.Mutex
	pendingIDs map[string]struct{}
	pendingMu  sync.Mutex
}

func NewWalletService(repo WalletRepository) *WalletService {
	return &WalletService{
		repo:       repo,
		pendingIDs: make(map[string]struct{}),
	}
}

func (s *WalletService) Transfer(ctx context.Context, fromID, toID string, amount Money) error {
	return s.TransferWithID(ctx, "", fromID, toID, amount)
}

func (s *WalletService) TransferWithID(ctx context.Context, transferID, fromID, toID string, amount Money) error {
	if transferID != "" {
		s.pendingMu.Lock()
		if _, exists := s.pendingIDs[transferID]; exists {
			s.pendingMu.Unlock()
			return errors.New("transfer already processed")
		}
		s.pendingIDs[transferID] = struct{}{}
		s.pendingMu.Unlock()
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	from, err := s.repo.GetWallet(ctx, fromID)
	if err != nil {
		return err
	}
	to, err := s.repo.GetWallet(ctx, toID)
	if err != nil {
		return err
	}

	if from.Balance.Naira() < amount.Naira() {
		return errors.New("insufficient funds")
	}

	from.Balance = Money(from.Balance.Naira() - amount.Naira())
	to.Balance = Money(to.Balance.Naira() + amount.Naira())

	if err := s.repo.SaveWallet(ctx, from); err != nil {
		return err
	}
	if err := s.repo.SaveWallet(ctx, to); err != nil {
		return err
	}

	return nil
}

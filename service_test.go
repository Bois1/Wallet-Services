package main

import (
	"context"
	"testing"
)

func TestWalletService_SuccessfulTransfer(t *testing.T) {
	repo := NewInMemoryWalletRepo()
	repo.SaveWallet(context.Background(), &Wallet{ID: "1", Owner: "Babatunde", Balance: Money(1000)})
	repo.SaveWallet(context.Background(), &Wallet{ID: "2", Owner: "Bimbo", Balance: Money(500)})

	service := NewWalletService(repo)

	err := service.Transfer(context.Background(), "1", "2", Money(200))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	w1, _ := repo.GetWallet(context.Background(), "1")
	w2, _ := repo.GetWallet(context.Background(), "2")

	if w1.Balance.Naira() != 800 {
		t.Errorf("Babatunde balance = %d, want 800", w1.Balance.Naira())
	}
	if w2.Balance.Naira() != 700 {
		t.Errorf("Bimbo balance = %d, want 700", w2.Balance.Naira())
	}
}

func TestWalletService_InsufficientFunds(t *testing.T) {
	repo := NewInMemoryWalletRepo()
	repo.SaveWallet(context.Background(), &Wallet{ID: "1", Owner: "Babatunde", Balance: Money(100)})

	service := NewWalletService(repo)

	err := service.Transfer(context.Background(), "1", "2", Money(200))
	if err == nil {
		t.Fatal("expected error, got none")
	}
	if err.Error() != "insufficient funds" {
		t.Errorf("wrong error: %v", err)
	}
}

func TestWalletService_Idempotency(t *testing.T) {
	repo := NewInMemoryWalletRepo()
	repo.SaveWallet(context.Background(), &Wallet{ID: "1", Owner: "Babatunde", Balance: Money(1000)})
	repo.SaveWallet(context.Background(), &Wallet{ID: "2", Owner: "Bimbo", Balance: Money(0)})

	service := NewWalletService(repo)

	err := service.TransferWithID(context.Background(), "tx-123", "1", "2", Money(100))
	if err != nil {
		t.Fatalf("first transfer failed: %v", err)
	}

	err = service.TransferWithID(context.Background(), "tx-123", "1", "2", Money(100))
	if err == nil {
		t.Fatal("expected idempotency error")
	}
	if err.Error() != "transfer already processed" {
		t.Errorf("wrong idempotency error: %v", err)
	}
}

func TestWalletService_WithFailingRepo(t *testing.T) {
	repo := FailingWalletRepo{}
	service := NewWalletService(repo)

	err := service.Transfer(context.Background(), "1", "2", Money(100))
	if err == nil {
		t.Fatal("expected error from failing repo")
	}
}

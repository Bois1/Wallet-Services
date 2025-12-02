package main

import (
	"context"
	"log"
)

func main() {

	repo := NewInMemoryWalletRepo()
	repo.SaveWallet(context.Background(), &Wallet{ID: "A", Owner: "Babatunde", Balance: Money(1000)})
	repo.SaveWallet(context.Background(), &Wallet{ID: "B", Owner: "Bimbo", Balance: Money(500)})

	service := NewWalletService(repo)

	err := service.TransferWithID(context.Background(), "tx-1", "A", "B", Money(300))
	if err != nil {
		log.Fatal("Transfer failed:", err)
	}

	a, _ := repo.GetWallet(context.Background(), "A")
	b, _ := repo.GetWallet(context.Background(), "B")

	log.Printf("Babatunde's balance: %s", a.Balance)
	log.Printf("Bimbo's balance: %s", b.Balance)
}

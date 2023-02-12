package bank_test

import (
	"fmt"
	"testing"

	bank "gopl.io/ch9/ex9_1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	bank.Deposit(200)

	if result := bank.WithDraw(100); !result {
		t.Error("WithDraw failed")
	}
	if got := bank.Balance(); got != 100 {
		t.Errorf("Remaining amount != Balance amount. Balance = %d", got)
	}
	if result := bank.WithDraw(150); result {
		t.Error("WithDraw should fail")
	}
}

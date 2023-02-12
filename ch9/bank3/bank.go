package bank

import "sync"

var (
	mu      sync.RWMutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

func Balance() int {
	mu.RLock() // readers lock
	defer mu.RUnlock()
	return balance
}

func WithDraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

// This function requires that the lock be held.
func deposit(amount int) { balance += amount }

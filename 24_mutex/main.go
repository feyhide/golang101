// MUTEX : race condition occurs when multiple
// process modifying the same resources
// and its not atomic enough

package main

import (
	"fmt"
	"sync"
)

// WITHOUT MUTEX
// type BankAccount struct {
// 	balance int
// }

// WITH MUTEX
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

// WITHOUT MUTEX
// func (acc *BankAccount) Deposit(amount int, w *sync.WaitGroup) {
// 	defer w.Done()
// 	acc.balance += amount
// }

// WITH MUTEX
func (acc *BankAccount) Deposit(amount int, w *sync.WaitGroup) {
	defer func() {
		w.Done()
		acc.mu.Unlock()
	}()

	acc.mu.Lock()
	acc.balance += amount

}

func main() {
	var wg sync.WaitGroup
	account := BankAccount{balance: 0}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go account.Deposit(100, &wg) // Multiple goroutines modifying balance
	}

	wg.Wait()
	fmt.Println("Final Balance:", account.balance)
	// im getting Final Balance: 9974200 instead of 10000000 without mutex

}

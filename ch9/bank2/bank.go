package bank

var deposits = make(chan int) // send amount to deposit
var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance += amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

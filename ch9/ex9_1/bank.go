package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdraws = make(chan withDrawInfo)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int { return <-balances }

type withDrawInfo struct {
	amount int
	result chan bool
}

func WithDraw(amount int) bool {
	result := make(chan bool)
	withdraws <- withDrawInfo{amount: amount, result: result}
	return <-result
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case info := <-withdraws:
			if info.amount > balance {
				info.result <- false
			} else {
				balance -= info.amount
				info.result <- true
			}
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

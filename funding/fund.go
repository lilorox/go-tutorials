package funding

type Fund struct {
    // private because lowercase
    balance int
}

// Function returning a pointer to a Fund struct
func NewFund(initialBalance int) *Fund {
    return &Fund{
        balance: initialBalance,
    }
}

// Method with a receiver returning the balance as int
func (f *Fund) Balance() int {
    return f.balance
}

func (f *Fund) Withdraw(amount int) {
    f.balance -= amount
}

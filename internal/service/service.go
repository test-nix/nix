package service

const (
	HelpCommand     = "/help"
	ExitCommand     = "0"
	TransactCommand = "1"
	BalanceCommand  = "2"
)

func ProcessTransaction(user *User, amount int64) error {
	if amount == 0 {
		return ErrZeroAmount
	}
	if amount < 0 && user.Balance+amount < 0 {
		return ErrInsufficientFunds
	}

	user.Balance += amount
	return nil
}

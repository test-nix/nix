package service

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

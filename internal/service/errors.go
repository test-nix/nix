package service

import "errors"

var ErrZeroAmount = errors.New("сума не може бути рівна нулю")
var ErrInsufficientFunds = errors.New("недостатньо коштів")

const (
	HelpCommand     = "/help"
	ExitCommand     = "0"
	TransactCommand = "1"
	BalanceCommand  = "2"
)

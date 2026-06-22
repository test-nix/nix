package service

import "errors"

var ErrZeroAmount = errors.New("the sum cannot be zero")
var ErrInsufficientFunds = errors.New("insufficient funds")

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"example/internal/service"
	"example/internal/service/helper"
	"example/internal/storage"
	"example/internal/transport"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome!")

	user, err := transport.LoginOrRegister(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hello, %s!\n", user.FirstName)

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Deposit / withdraw funds")
		fmt.Println("2. View balance")
		fmt.Println("/help — help  |  0 — Exit")

		choice, err := transport.ReadInput(reader, "Select action: ")
		if err != nil {
			fmt.Println("Reading error:", err)
			continue
		}

		if choice == service.HelpCommand {
			helper.PrintHelp()
			continue
		}
		if choice == service.ExitCommand {
			fmt.Println("Goodbye!")
			break
		}

		switch choice {
		case service.TransactCommand:
			input, _ := transport.ReadInput(reader, "Enter amount in dollars: ")
			amount, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Error: entered not a number.")
				continue
			}
			err = service.ProcessTransaction(user, amount)
			if errors.Is(err, service.ErrInsufficientFunds) {
				fmt.Printf("✗ Insufficient funds! Balance: $%d\n", user.Balance)
			} else if errors.Is(err, service.ErrZeroAmount) {
				fmt.Println("✗ Enter a non-zero sum")
			} else {
				err = storage.UpdateBalance(user)
				if err != nil {
					fmt.Println("Saving error:", err)
				} else {
					fmt.Printf("✓ Success. New balance: $%d\n", user.Balance)
				}
			}

		case service.BalanceCommand:
			helper.PrintUserInfo(user)

		default:
			fmt.Println("Unknown command. Enter /help for help.")
		}
	}
}

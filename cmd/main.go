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

func loginOrRegister(reader *bufio.Reader) (*service.User, error) {
	fmt.Println("\n1. Увійти")
	fmt.Println("2. Зареєструватись")

	choice, _ := transport.ReadInput(reader, "Виберіть: ")
	switch choice {
	case "1":
		return login(reader)
	case "2":
		return register(reader)
	default:
		return nil, fmt.Errorf("невідома команда")
	}
}

func login(reader *bufio.Reader) (*service.User, error) {
	username, _ := transport.ReadInput(reader, "Username: ")
	password, _ := transport.ReadInput(reader, "Пароль: ")

	user, err := storage.LoadUser(username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, fmt.Errorf("невірний пароль")
	}
	return user, nil
}

func register(reader *bufio.Reader) (*service.User, error) {
	username, _ := transport.ReadInput(reader, "Введіть username: ")

	_, err := storage.LoadUser(username)
	if err == nil {
		return nil, fmt.Errorf("користувач '%s' вже існує", username)
	}

	firstname, _ := transport.ReadInput(reader, "Введіть імʼя: ")
	lastname, _ := transport.ReadInput(reader, "Введіть прізвище: ")
	password, _ := transport.ReadInput(reader, "Придумайте пароль: ")

	user := &service.User{
		ID:        storage.GetNextID(),
		Username:  username,
		FirstName: firstname,
		LastName:  lastname,
		Balance:   0,
		Password:  password,
	}

	err = storage.RegisterUser(user)
	if err != nil {
		return nil, fmt.Errorf("помилка реєстрації: %w", err)
	}

	fmt.Println("✓ Реєстрація успішна!")
	return user, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ласкаво просимо!")

	user, err := loginOrRegister(reader)
	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}

	fmt.Printf("Привіт, %s!\n", user.FirstName)

	for {
		fmt.Println("\n--- Меню ---")
		fmt.Println("1. Поповнити / зняти кошти")
		fmt.Println("2. Переглянути баланс")
		fmt.Println("/help — довідка  |  0 — Вихід")

		choice, err := transport.ReadInput(reader, "Виберіть дію: ")
		if err != nil {
			fmt.Println("Помилка читання:", err)
			continue
		}

		if choice == service.HelpCommand {
			helper.PrintHelp()
			continue
		}
		if choice == service.ExitCommand {
			fmt.Println("До побачення!")
			break
		}

		switch choice {
		case service.TransactCommand:
			input, _ := transport.ReadInput(reader, "Введіть суму в доларах: ")
			amount, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Помилка: введено не число.")
				continue
			}
			err = service.ProcessTransaction(user, amount)
			if errors.Is(err, service.ErrInsufficientFunds) {
				fmt.Printf("✗ Недостатньо коштів! Баланс: $%d\n", user.Balance)
			} else if errors.Is(err, service.ErrZeroAmount) {
				fmt.Println("✗ Введіть ненульову суму")
			} else {
				err = storage.UpdateBalance(user)
				if err != nil {
					fmt.Println("Помилка збереження:", err)
				} else {
					fmt.Printf("✓ Успішно. Новий баланс: $%d\n", user.Balance)
				}
			}

		case service.BalanceCommand:
			helper.PrintUserInfo(user)

		default:
			fmt.Println("Невідома команда. Введіть /help для довідки.")
		}
	}
}

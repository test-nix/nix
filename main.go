package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID        int64
	Username  string
	FirstName string
	LastName  string
	Balance   int64
}

var ErrZeroAmount = errors.New("сума не може бути рівна нулю")
var ErrInsufficientFunds = errors.New("недостатньо коштів")

const (
	HelpCommand     = "/help"
	ExitCommand     = "0"
	TransactCommand = "1"
	BalanceCommand  = "2"
)

func printHelp() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║         ДОВІДКА ПО ПРОГРАМІ          ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  Команди меню:                       ║")
	fmt.Println("║  1 — Поповнити або зняти кошти       ║")
	fmt.Println("║  2 — Переглянути баланс              ║")
	fmt.Println("║  /help — Показати цю довідку         ║")
	fmt.Println("║  0 — Вийти з програми                ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  Як вводити суми:                    ║")
	fmt.Println("║  2000 або +2000  — поповнення        ║")
	fmt.Println("║  -2000           — зняття            ║")
	fmt.Println("╚══════════════════════════════════════╝")
}

func printUserInfo(user User) {
	fmt.Println("┌─────────────────────────────────┐")
	fmt.Printf("│  ID         : %d\n", user.ID)
	fmt.Printf("│  Username   : %s\n", user.Username)
	fmt.Printf("│  Імʼя       : %s\n", user.FirstName)
	fmt.Printf("│  Прізвище   : %s\n", user.LastName)
	fmt.Printf("│  Баланс     : $%d\n", user.Balance)
	fmt.Println("└─────────────────────────────────┘")
}

func processTransaction(user *User, amount int64) error {
	if amount == 0 {
		return ErrZeroAmount
	}
	if amount < 0 && user.Balance+amount < 0 {
		return ErrInsufficientFunds
	}

	user.Balance += amount
	return nil
}

func readInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hello, World!")

	users := map[string]*User{
		"tom_99":  {ID: 1, Username: "tom_99", FirstName: "Tom", LastName: "Hollands", Balance: 1000},
		"nik_22":  {ID: 2, Username: "nik_22", FirstName: "Nik", LastName: "Luks", Balance: 2000},
		"vlad_07": {ID: 3, Username: "vlad_07", FirstName: "Vlad", LastName: "Kots", Balance: 3000},
	}

	fmt.Println("Ласкаво просимо! Введіть /help для довідки.")

	for {
		fmt.Println("\n--- Меню ---")
		fmt.Println("1. Поповнити / зняти кошти")
		fmt.Println("2. Переглянути баланс")
		fmt.Println("/help — довідка  |  0 — Вихід")

		choice, err := readInput(reader, "Виберіть дію: ")
		if err != nil {
			fmt.Println("Помилка читання:", err)
			continue
		}

		if choice == HelpCommand {
			printHelp()
			continue
		}

		if choice == ExitCommand {
			fmt.Println("До побачення!")
			break
		}

		username, err := readInput(reader, "Введіть username: ")
		if err != nil {
			fmt.Println("Помилка читання:", err)
			continue
		}

		user, exists := users[username]
		if !exists {
			fmt.Printf("Користувача '%s' не знайдено.\n", username)
			continue
		}

		switch choice {
		case TransactCommand:
			input, _ := readInput(reader, "Введіть суму в доларах: ")
			amount, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Помилка: введено не число.")
				continue
			}
			err = processTransaction(user, amount)
			if errors.Is(err, ErrInsufficientFunds) {
				fmt.Printf("✗ Недостатньо коштів! Баланс: $%d\n", user.Balance)
			} else if errors.Is(err, ErrZeroAmount) {
				fmt.Println("✗ Введіть ненульову суму")
			} else {
				fmt.Printf("✓ Успішно. Новий баланс: $%d\n", user.Balance)
			}
		case BalanceCommand:
			printUserInfo(*user)

		default:
			fmt.Println("Невідома команда. Введіть /help для довідки.")
		}
	}
}

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
	Password  string
}

var ErrZeroAmount = errors.New("the sum cannot be zero")
var ErrInsufficientFunds = errors.New("insufficient funds")

const (
	HelpCommand     = "/help"
	ExitCommand     = "0"
	TransactCommand = "1"
	BalanceCommand  = "2"
)

func printHelp() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║         PROGRAM INFORMATION          ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  Menu commands:                      ║")
	fmt.Println("║  1 — Deposit or withdraw funds       ║")
	fmt.Println("║  2 — View balance           		    ║")
	fmt.Println("║  /help — Show this help         	    ║")
	fmt.Println("║  0 — Exit the program                ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  How to enter amounts:               ║")
	fmt.Println("║  2000 or +2000  — deposit             ║")
	fmt.Println("║  -2000           — withdrawal          ║")
	fmt.Println("╚══════════════════════════════════════╝")
}

func printUserInfo(user User) {
	fmt.Println("┌─────────────────────────────────┐")
	fmt.Printf("│  ID         : %d\n", user.ID)
	fmt.Printf("│  Username   : %s\n", user.Username)
	fmt.Printf("│  First Name : %s\n", user.FirstName)
	fmt.Printf("│  Last Name  : %s\n", user.LastName)
	fmt.Printf("│  Balance    : $%d\n", user.Balance)
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

const filePath = "users.txt"

func loadUser(username string) (*User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("user file not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		user, err := parseLine(scanner.Text())
		if err != nil {
			continue
		}
		if user.Username == username {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user '%s' not found", username)
}

func updateBalance(user *User) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 6 && parts[1] == user.Username {
			line = fmt.Sprintf("%d|%s|%s|%s|%d|%s",
				user.ID, user.Username, user.FirstName,
				user.LastName, user.Balance, user.Password)
		}
		lines = append(lines, line)
	}
	file.Close()

	file, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func registerUser(user *User) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	line := fmt.Sprintf("%d|%s|%s|%s|%d|%s\n",
		user.ID, user.Username, user.FirstName,
		user.LastName, user.Balance, user.Password)

	_, err = file.WriteString(line)
	return err
}

func getNextID() int64 {
	file, err := os.Open(filePath)
	if err != nil {
		return 1
	}
	defer file.Close()

	var maxID int64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		if len(parts) >= 1 {
			id, err := strconv.ParseInt(parts[0], 10, 64)
			if err == nil && id > maxID {
				maxID = id
			}
		}
	}
	return maxID + 1
}

func parseLine(line string) (*User, error) {
	parts := strings.Split(line, "|")
	if len(parts) != 6 {
		return nil, fmt.Errorf("invalid line format")
	}

	id, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}
	balance, err := strconv.ParseInt(parts[4], 10, 64)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id,
		Username:  parts[1],
		FirstName: parts[2],
		LastName:  parts[3],
		Balance:   balance,
		Password:  parts[5],
	}, nil
}

func loginOrRegister(reader *bufio.Reader) (*User, error) {
	fmt.Println("\n--- Enter ---")
	fmt.Println("1. Login")
	fmt.Println("2. Register")

	choice, _ := readInput(reader, "Select: ")

	switch choice {
	case "1":
		return login(reader)
	case "2":
		return register(reader)
	default:
		return nil, fmt.Errorf("unknown command")
	}
}

func login(reader *bufio.Reader) (*User, error) {
	username, _ := readInput(reader, "Username: ")
	password, _ := readInput(reader, "Password: ")

	user, err := loadUser(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, fmt.Errorf("incorrect password")
	}

	return user, nil
}

func register(reader *bufio.Reader) (*User, error) {
	username, _ := readInput(reader, "Enter username: ")

	_, err := loadUser(username)
	if err == nil {
		return nil, fmt.Errorf("user '%s' already exists", username)
	}

	firstname, _ := readInput(reader, "Enter first name: ")
	lastname, _ := readInput(reader, "Enter last name: ")
	password, _ := readInput(reader, "Enter password: ")

	user := &User{
		ID:        getNextID(),
		Username:  username,
		FirstName: firstname,
		LastName:  lastname,
		Balance:   0,
		Password:  password,
	}

	err = registerUser(user)
	if err != nil {
		return nil, fmt.Errorf("registration error: %w", err)
	}

	fmt.Println("✓ Registration successful!")
	return user, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	users := map[string]*User{
		"tom_99":  {ID: 1, Username: "tom_99", FirstName: "Tom", LastName: "Hollands", Balance: 1000},
		"nik_22":  {ID: 2, Username: "nik_22", FirstName: "Nik", LastName: "Luks", Balance: 2000},
		"vlad_07": {ID: 3, Username: "vlad_07", FirstName: "Vlad", LastName: "Kots", Balance: 3000},
	}

	fmt.Println("Welcome! Enter /help for help.")

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Deposit / withdraw funds")
		fmt.Println("2. View balance")
		fmt.Println("/help — help  |  0 — Exit")

		choice, err := readInput(reader, "Select action: ")
		if err != nil {
			fmt.Println("Reading error:", err)
			continue
		}

		if choice == HelpCommand {
			printHelp()
			continue
		}

		if choice == ExitCommand {
			fmt.Println("Goodbye!")
			break
		}

		username, err := readInput(reader, "Enter username: ")
		if err != nil {
			fmt.Println("Reading error:", err)
			continue
		}

		user, exists := users[username]
		if !exists {
			fmt.Printf("User '%s' not found.\n", username)
			continue
		}

		switch choice {
		case TransactCommand:
			input, _ := readInput(reader, "Enter amount in dollars: ")
			amount, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Error: entered not a number.")
				continue
			}
			err = processTransaction(user, amount)
			if errors.Is(err, ErrInsufficientFunds) {
				fmt.Printf("✗ Insufficient funds! Balance: $%d\n", user.Balance)
			} else if errors.Is(err, ErrZeroAmount) {
				fmt.Println("✗ Enter a non-zero sum")
			} else {
				fmt.Printf("✓ Success. New balance: $%d\n", user.Balance)
			}
		case BalanceCommand:
			printUserInfo(*user)

		default:
			fmt.Println("Unknown command. Enter /help for help.")
		}
	}
}

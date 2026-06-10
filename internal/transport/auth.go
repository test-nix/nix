package transport

import (
	"bufio"
	"fmt"

	"example/internal/service"
	"example/internal/storage"
)

func LoginOrRegister(reader *bufio.Reader) (*service.User, error) {
	fmt.Println("\n1. Sign in")
	fmt.Println("2. Register")

	choice, _ := ReadInput(reader, "Select: ")
	switch choice {
	case "1":
		return Login(reader)
	case "2":
		return Register(reader)
	default:
		return nil, fmt.Errorf("unknown command")
	}
}

func Login(reader *bufio.Reader) (*service.User, error) {
	username, _ := ReadInput(reader, "Username: ")
	password, _ := ReadInput(reader, "Password: ")

	user, err := storage.LoadUser(username)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, fmt.Errorf("incorrect password")
	}
	return user, nil
}

func Register(reader *bufio.Reader) (*service.User, error) {
	username, _ := ReadInput(reader, "Enter username: ")

	_, err := storage.LoadUser(username)
	if err == nil {
		return nil, fmt.Errorf("user '%s' already exists", username)
	}

	firstname, _ := ReadInput(reader, "Enter first name: ")
	lastname, _ := ReadInput(reader, "Enter last name: ")
	password, _ := ReadInput(reader, "Enter password: ")

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
		return nil, fmt.Errorf("registration error: %w", err)
	}

	fmt.Println("✓ Registration successful!")
	return user, nil
}

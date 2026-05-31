package storage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example/internal/service"
)

const filePath = "users.txt"

func LoadUser(username string) (*service.User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("файл користувачів не знайдено")
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
	return nil, fmt.Errorf("користувача '%s' не знайдено", username)
}

func UpdateBalance(user *service.User) error {
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

func RegisterUser(user *service.User) error {
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

func GetNextID() int64 {
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

func parseLine(line string) (*service.User, error) {
	parts := strings.Split(line, "|")
	if len(parts) != 6 {
		return nil, fmt.Errorf("невірний формат")
	}
	id, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}
	balance, err := strconv.ParseInt(parts[4], 10, 64)
	if err != nil {
		return nil, err
	}
	return &service.User{
		ID:        id,
		Username:  parts[1],
		FirstName: parts[2],
		LastName:  parts[3],
		Balance:   balance,
		Password:  parts[5],
	}, nil
}

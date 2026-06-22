package helper

import (
	"example/internal/service"
	"fmt"
)

func PrintHelp() {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║         PROGRAM INFORMATION          ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  Menu commands:                      ║")
	fmt.Println("║  1 — Deposit or withdraw funds       ║")
	fmt.Println("║  2 — View balance                     ║")
	fmt.Println("║  /help — Show this help               ║")
	fmt.Println("║  0 — Exit the program                  ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Println("║  How to enter amounts:               ║")
	fmt.Println("║  2000 or +2000  — deposit             ║")
	fmt.Println("║  -2000           — withdrawal          ║")
	fmt.Println("╚══════════════════════════════════════╝")
}

func PrintUserInfo(user *service.User) {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║       USER INFORMATION              ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Printf("║  ID         : %d\n", user.ID)
	fmt.Printf("║  Username   : %s\n", user.Username)
	fmt.Printf("║  First Name : %s\n", user.FirstName)
	fmt.Printf("║  Last Name  : %s\n", user.LastName)
	fmt.Printf("║  Balance    : $%d\n", user.Balance)
	fmt.Println("╚══════════════════════════════════════╝")
}

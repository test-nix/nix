package helper

import (
	"example/internal/service"
	"fmt"
)

func PrintHelp() {
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

func PrintUserInfo(user *service.User) {
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║       ІНФОРМАЦІЯ ПРО КОРИСТУВАЧА     ║")
	fmt.Println("╠══════════════════════════════════════╣")
	fmt.Printf("║  ID         : %d\n", user.ID)
	fmt.Printf("║  Username   : %s\n", user.Username)
	fmt.Printf("║  Імʼя       : %s\n", user.FirstName)
	fmt.Printf("║  Прізвище   : %s\n", user.LastName)
	fmt.Printf("║  Баланс     : $%d\n", user.Balance)
	fmt.Println("╚══════════════════════════════════════╝")
}

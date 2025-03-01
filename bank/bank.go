package main

import (
	"fmt"
	"os"
)

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositeAmount float64
			fmt.Scan(&depositeAmount)

			if depositeAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositeAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
			// break
		}

		/*

			if choice == 1 {
				fmt.Println("Your balance is", accountBalance)
			} else if choice == 2 {
				fmt.Print("Your deposit: ")
				var depositeAmount float64
				fmt.Scan(&depositeAmount)

				if depositeAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					// return
					continue
				}

				accountBalance += depositeAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
			} else if choice == 3 {
				fmt.Print("Withdrawal amount: ")
				var withdrawalAmount float64
				fmt.Scan(&withdrawalAmount)

				if withdrawalAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					return
				}

				if withdrawalAmount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have.")
					return
				}

				accountBalance -= withdrawalAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
			} else {
				fmt.Println("Goodbye!")
				// return
				break
			}

		*/

	}

}

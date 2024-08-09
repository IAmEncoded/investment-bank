package main

import (
	"fmt"

	"example.com/investment-bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic(err)
	}
	
	var depositAmount float64
	var withdrawlAmmount float64

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Balance is: ", accountBalance)
		case 2:
			fmt.Print("Input the amount you want to deposit: ")
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
	
			accountBalance += depositAmount
	
			fmt.Println("Updated balance:", accountBalance)

			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawl ammount: ")
			fmt.Scan(&withdrawlAmmount)
			
			if withdrawlAmmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
	
			if withdrawlAmmount > accountBalance {
				fmt.Println("You cannot withdraw more money than is in your account.")
				continue
			}
	
			accountBalance -= withdrawlAmmount
			
			fmt.Println("Updated balance:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}
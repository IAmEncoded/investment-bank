package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err= getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic(err)
	}
	
	var depositAmount float64
	var withdrawlAmmount float64

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

			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
		}
	}
}

func writeBalanceToFile(balance float64) {
	balanceTxt := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceTxt), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to find balance file")
	}

	balanceTxt := string(data)
	balance, err := strconv.ParseFloat(balanceTxt, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}
package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "./balance.txt"

func main() {
	var accountBalance, error = fileops.GetFloatFromFile(accountBalanceFile)
	if error != nil {
		fmt.Println("ERROR")
		fmt.Println(error)
		fmt.Println("------------")
		panic(error)
	}
	fmt.Println("Welcome to the Go Bank")
	fmt.Println(randomdata.PhoneNumber())
	for {
		presentOptions()
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("Your Balance : ", accountBalance)
		case 2:
			fmt.Print("Money to deposit : ")
			var moneyToDeposit float64
			fmt.Scan(&moneyToDeposit)
			if moneyToDeposit <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance += moneyToDeposit
			fmt.Printf("%v deposited.\n", moneyToDeposit)
			fmt.Printf("Final Balance : %v\n", accountBalance)
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		case 3:
			fmt.Print("Money to Withdraw : ")
			var moneyToWithdraw float64
			fmt.Scan(&moneyToWithdraw)
			if moneyToWithdraw > accountBalance {
				fmt.Println("Withdraw amount is more than account balance")
				continue
			}
			if moneyToWithdraw <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0")
				continue
			}
			accountBalance -= moneyToWithdraw
			fmt.Printf("%v debited from account\n", moneyToWithdraw)
			fmt.Printf("Final Balance : %v\n", accountBalance)
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		default:
			println("Bye !")
			return
		}
	}
}

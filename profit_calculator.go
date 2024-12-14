package main

import (
	"errors"
	"fmt"
	"os"
)

func writeResultsToFile(revenue, expenses, taxRate, ebt, profit, ratio float64) {
	data := fmt.Sprintf("Revenue : %.2f , Expenses : %.2f , taxRate : %.2f%% , Earnings before Tax = %.3f , Profit = %.2f , ratio = %.2f\n", revenue, expenses, taxRate, ebt, profit, ratio)
	os.WriteFile("./records.txt", []byte(data), os.ModeAppend)
}

func main() {
	revenue, err1 := getUserInput("Revenue : ")
	expenses, err2 := getUserInput("Expenses : ")
	taxRate, err3 := getUserInput("Tax Rate : ")
	if err1 != nil || err2 != nil || err3 != nil {
		panic(err1)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax : %.2f\nEarnings after tax(Profit) : %.2f\nEBT / Profit : %.3f\n", ebt, profit, ratio)

	writeResultsToFile(revenue, expenses, taxRate, ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - revenue*(taxRate/100)
	ratio = ebt / profit
	return
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0.0, errors.New("value can't be negative or zero")
	}
	return userInput, nil
}

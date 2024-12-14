package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 6.0

func main1() {
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Investment Amount : ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Investment Time : ")
	fmt.Scan(&years)
	fmt.Print("Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	// var futureValue float64 = calculateFutureValue(investmentAmount, expectedReturnRate, years)
	// var futureRealValue float64 = calculateRealFutureValue(futureValue, inflationRate, years)

	var futureValue, futureRealValue = calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Absolute Future Value : %.2f\n", futureValue)
	fmt.Printf("Future Value after Inflation adjustment : %.2f\n", futureRealValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}

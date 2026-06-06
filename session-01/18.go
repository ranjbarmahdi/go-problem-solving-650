// package main

// import "fmt"

// const (
// 	annualPenConsumption = 50
// 	annualA4Consumption  = 150
// )

// func main() {
// 	var penCurrentPrice, a4CurrentPrice, inflationRate float64

// 	fmt.Print("Enter A4 Current Price: ")
// 	fmt.Scan(&a4CurrentPrice)

// 	fmt.Print("Enter Pen Current Price: ")
// 	fmt.Scan(&penCurrentPrice)

// 	fmt.Print("Enter Next Year Inflation Rate: ")
// 	fmt.Scan(&inflationRate)

// 	growthInExpenses := ((a4CurrentPrice * annualA4Consumption) + (penCurrentPrice * annualPenConsumption)) * inflationRate / 100
// 	fmt.Print("Growth In Company Expenses:", growthInExpenses)
// }

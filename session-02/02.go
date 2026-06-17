// package main

// import "fmt"

// func main() {
// 	for {
// 		var number int

// 		fmt.Print("Enter number: ")
// 		fmt.Scan(&number)

// 		var sum int
// 		for i := 1; i < number; i++ {
// 			if number%i == 0 {
// 				sum += i
// 			}
// 		}

// 		if sum == number {
// 			fmt.Println("Number ", number, " is perfect.")
// 		} else {
// 			fmt.Println("Number ", number, " is not perfect.")
// 		}

// 		var continueProgress string
// 		fmt.Print("Do you want to continue(y / n): ")
// 		fmt.Scan(&continueProgress)

// 		if continueProgress == "n" {
// 			fmt.Println("End")
// 			break
// 		}
// 	}
// }

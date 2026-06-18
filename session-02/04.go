// package main

// import "fmt"

// const (
// 	TaxForMoreThan400000 = 0.1
// 	TaxForMoreThan500000 = 0.15
// 	TaxForMoreThan700000 = 0.1
// )

// func main() {
// 	var numberOfEmployees int

// 	fmt.Print("Enter number of employees: ")
// 	fmt.Scan(&numberOfEmployees)

// 	var (
// 		maxSalary float64
// 		maxEmpId  string
// 	)

// 	for range numberOfEmployees {
// 		var (
// 			empId  string
// 			salary float64
// 		)

// 		fmt.Print("Enter Employee Id: ")
// 		fmt.Scan(&empId)

// 		fmt.Print("Enter Salary: ")
// 		fmt.Scan(&salary)

// 		switch {
// 		case salary > 400000:
// 			salary = salary - salary*TaxForMoreThan400000
// 		case salary > 500000:
// 			salary = salary - salary*TaxForMoreThan500000
// 		case salary > 700000:
// 			salary = salary - salary*TaxForMoreThan700000
// 		}

// 		fmt.Printf("Net Pay For User With %s Employee Id is %0.3f Rials.\n\n", empId, salary)

// 		if salary > maxSalary {
// 			maxSalary = salary
// 			maxEmpId = empId
// 		}
// 	}

// 	fmt.Printf("Maximum Net Pay Is For User With %s Employee Id, %0.3f Rials.", maxEmpId, maxSalary)

// }

// package main

// import "fmt"

// func main() {
// 	var numberOfStudents int

// 	fmt.Print("Enter Number Of Students: ")
// 	fmt.Scan(&numberOfStudents)

// 	var (
// 		max1, max2, avg float64
// 		id, id1, id2    string
// 	)
// 	for range numberOfStudents {
// 		fmt.Print("Enter average: ")
// 		fmt.Scan(&avg)

// 		fmt.Print("Enter id: ")
// 		fmt.Scan(&id)

// 		if avg > max1 {
// 			max2 = max1
// 			id2 = id1

// 			max1 = avg
// 			id1 = id
// 		} else if avg > max2 {
// 			max2 = avg
// 			id2 = id
// 		}
// 	}

// 	fmt.Println(id1)
// 	fmt.Println(max1)
// 	fmt.Println(id2)
// 	fmt.Println(max2)

// }

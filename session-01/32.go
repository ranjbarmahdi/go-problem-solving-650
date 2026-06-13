// package main

// import "fmt"

// const (
// 	NumberOfPageLines = int64(30)
// 	LineSizeInBytes   = int64(80)
// 	BytesPerGB        = int64(1024 * 1024 * 1024)
// )

// func main() {
// 	var pageCount int64
// 	var flashCapacityInGb int64

// 	fmt.Print("Enter number of book pages: ")
// 	fmt.Scan(&pageCount)

// 	fmt.Print("Enter capacity of usb flash in GB: ")
// 	fmt.Scan(&flashCapacityInGb)

// 	bookSizeBytes := pageCount * NumberOfPageLines * LineSizeInBytes
// 	flashSizeBytes := flashCapacityInGb * BytesPerGB

// 	booksToStore := flashSizeBytes / bookSizeBytes

// 	fmt.Println("Book size in bytes:", bookSizeBytes)
// 	fmt.Println("Flash size in bytes:", flashSizeBytes)
// 	fmt.Println("Books can be stored:", booksToStore)
// }

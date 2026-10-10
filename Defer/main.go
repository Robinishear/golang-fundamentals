package main

import "fmt"

// func deferred(result int) {
// 	fmt.Println("i am from defer function ", result) // defer function executes late (at the end of the surrounding function)

// }

func deferFunction() int {

	result := 10

	defer func() {
		fmt.Println("i am from defer function ", result)
	}()

	// defer deferred(result) // defer function executes late (at the end of the surrounding function)

	fmt.Println("i am from defer deferFunction ", result)

	result += 100

	return result
}

func main() {
	fmt.Println("return result", deferFunction())

	// defer fmt.Println("i am from defer function ") // defer function executes late (at the end of the surrounding function)

	//  fmt.Println("i am from main function ")

}

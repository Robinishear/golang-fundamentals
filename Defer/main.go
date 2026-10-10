package main

import "fmt"

func deferred(result int) {
	fmt.Println("i am from defer function ", result) // defer function executes late (at the end of the surrounding function)

}

func deferFunction() int {
	result := 10

	defer deferred(result) // defer function executes late (at the end of the surrounding function)

	fmt.Println("i am from defer deferFunction ", result) // defer function executes late (at the end of the surrounding function)
	return result
}

func main() {
	deferFunction()

	// defer fmt.Println("i am from defer function ") // defer function executes late (at the end of the surrounding function)

	//  fmt.Println("i am from main function ")

}

// package main

// import "fmt"

// // func deferred(result int) {
// // 	fmt.Println("i am from defer function ", result) // defer function executes late (at the end of the surrounding function)

// // }

// func deferFunction() int {

// 	result := 10

// 	defer func() {
// 		result += 200
// 		fmt.Println("i am from defer function ", result)
// 	}()

// 	// defer deferred(result) // defer function executes late (at the end of the surrounding function)

// 	fmt.Println("i am from defer deferFunction ", result)

// 	result += 100

// 	return result
// }

// func main() {
// 	fmt.Println("return result", deferFunction())

// 	// defer fmt.Println("i am from defer function ") // defer function executes late (at the end of the surrounding function)

// 	//  fmt.Println("i am from main function ")

// }

// i am from defer deferFunction  10
// i am from defer function  310
// return result 110

//====================
// name return value
// i am from defer deferFunction  10
// i am from defer function  310
// return result 310

package main

import "fmt"

func deferFunction() int {

	result := 10

	defer func() {
		result += 200
		fmt.Println("i am from defer function ", result)
	}()

	// defer deferred(result) // defer function executes late (at the end of the surrounding function)

	fmt.Println("i am from defer deferFunction ", result)

	result += 100

	return result
}

func nameExample() (result int) {

	result = 10

	defer func() {
		result += 200
		fmt.Println("i am from defer function ", result)
	}()
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	defer fmt.Println(5)

	fmt.Println("i am from defer deferFunction ", result)

	result += 100

	return
}

func main() {
	fmt.Println("name return result", nameExample())
	// fmt.Println("==============================")
	// fmt.Println("return result", deferFunction())

	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// defer fmt.Println(4)
	// defer fmt.Println(5)

}

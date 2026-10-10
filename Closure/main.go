package main

import "fmt"

// func multiplyBy(factor int) func(int) int {
// 	return func(x int) int {
// 		return x * factor
// 	}
// }

func makCounter() func() int {
	count := 0

	inner := func() int {
		count++
		return count
	}
	return inner
}

func main() {

	counter := makCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// multiplyByTwo := multiplyBy(2)
	// fmt.Println(multiplyByTwo(1000))
}

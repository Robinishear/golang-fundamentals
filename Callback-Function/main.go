package main

import "fmt"

func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

func main() {

	// greet := func() {
	// 	fmt.Println("Hello there")
	// }
	// process(greet)

	add := func(n1 int, n2 int) int {
		return n1 + n2
	}
	multiply := func(n1 int, n2 int) int {
		return n1 * n2
	}

	fmt.Println(calculate(10, 20, add))
	fmt.Println(calculate(10, 20, multiply))

	result := calculate(10, 20, func(x int, y int) int {
		return x - y
	})
	fmt.Println(result)
}

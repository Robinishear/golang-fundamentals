package main

import "fmt"

func add(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total = total + number
	}
	return total
}

func greet(prefix string, mps ...string) {
	for _, mp := range mps {
		fmt.Println(prefix, mp)
	}

}

func main() {
	sum := add(1, 2, 3, 4, 5)
	fmt.Println(sum)

	mps := []string{"John", "Doe", "Smith", "Jane", "Doe", "Smith", "Jane", "Doe", "Smith", "Jane", "Doe", "Smith", "Jane", "Doe", "Smith", "Jane", "Doe", "Smith", "Jane", "Doe", "Smith"}
	greet("welcome bro..?", mps...)
}

// flexible amount of argument
// must be the last parameter
// internally slice

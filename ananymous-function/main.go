package main

// import "fmt"

import "fmt"

func main() {

	// Ananymous Function

	coffeeOrder := func() {
		fmt.Println("Coffee order placed")
	}
	coffeeOrder()

	// Immediately Invoked Function Expression (IIFE)
	func(CoffeeType string) {
		fmt.Printf("Coffee order placed  %s.........", CoffeeType)

	}("Latte")

	//

	mackCoffee := func() {
		Coffee := "Black Coffee"
		price := "150"

		fmt.Printf("Mack Coffee Price %s TK %s", Coffee, price)
	}
	mackCoffee()

}

package main

import "fmt"

type Animal interface {
	speak()
} 

type Dog struct{}
type Cat struct{}
type Human struct{
	name string
}

// dog
func (d Dog) speak() {
	fmt.Println("Woof!!!")
}

// cat
func (c Cat) speak() {
	fmt.Println("Meow!!!")
}

// human
func (h Human) speak() {
	fmt.Println("Hallo My Name is",h.name)
}
// makeSound function
func makeSound(a Animal) {
	a.speak()
}


// main function
func main() {

	dexter := Dog{}
robin := Human{name: "Robin"}
bella := Cat{}
	makeSound(dexter)
		makeSound(bella)

		makeSound(robin)

}

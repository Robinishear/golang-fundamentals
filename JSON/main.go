package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"person_name"`  // struct tag
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {

	p := person{
		Name: "John Doe",
		Age:  30,
		City: "New York",
	}

	rowJson, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)

	}
	fmt.Println("JSON:", string(rowJson))
}


package main

import "fmt"

func Process(data any) {
	strData, ok := data.(string)
	if ok {
		fmt.Println(len(strData))

	}

	intData, ok := data.(int)
	if ok {
		fmt.Println(intData + 100)

	}
}

func main() {
	Process("100")
}

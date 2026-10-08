package test
import "fmt"


type MakPaymentMethod struct {
}

func (mk *MakPaymentMethod) Pay(amount float64) {
	fmt.Printf("Paying %.2f using MakPaymentMethod successfully\n", amount)
}
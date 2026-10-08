package main

import (
	"golang-fundamentals/payment"
	"golang-fundamentals/test"

	"github.com/fatih/color"
)

func main() {

	bkash := payment.NewBkash("25923DLO")
	paymentService1 := payment.NewPaymentService(bkash)
	paymentService1.Checkout()

	nagad := payment.NewNagad("NAGAD123")
	paymentService2 := payment.NewPaymentService(nagad)
	paymentService2.Checkout()

	mk := &test.MakPaymentMethod{}
	paymentService3 := payment.NewPaymentService(mk)
	paymentService3.Checkout()
	color.Cyan("Prints text in cyan.")
	color.RGB(255, 128, 0).Println("foreground orange")


}

package main



func main() {

	bkash := NewBkash("25923DLO")
	paymentService1 := NewPaymentService(bkash)
	paymentService1.checkout()

	nagad := NewNagad("NAGAD123")
	paymentService2 := NewPaymentService(nagad)
	paymentService2.checkout()

    mk := &MakPaymentMethod{}
	paymentService3 := NewPaymentService(mk)
	paymentService3.checkout()
}

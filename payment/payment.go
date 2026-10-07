package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}

type Bkash struct {
	apiKey string
}
type Nagad struct {
	apiKey string
}

func (bk *Bkash) pay(amount float64) {
	fmt.Printf("Paying %.2f using Bkash with API Key: %s\n", amount, bk.apiKey)
}

func (ng *Nagad) pay(amount float64) {
	fmt.Printf("Paying %.2f using Nagad with API Key: %s\n", amount, ng.apiKey)
}

type PaymentService struct{
	method PaymentMethod
}

func NewNagad (apiKey string) *Nagad {
	return &Nagad{
		apiKey: apiKey,
	}
}
func NewBkash (apiKey string) *Bkash {
	return &Bkash{
		apiKey: apiKey,
	}
}

func NewPaymentService(method PaymentMethod) *PaymentService {
	return &PaymentService{
		method: method,
	}
}

func (ps *PaymentService) checkout() {
	ps.method.pay(10000.000)
}

type MakPaymentMethod struct {
}

func (mk *MakPaymentMethod) pay(amount float64) {
	fmt.Printf("Paying %.2f using MakPaymentMethod successfully\n", amount)
}
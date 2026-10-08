package payment

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
}

type Bkash struct {
	apiKey string
}
type Nagad struct {
	apiKey string
}

func (bk *Bkash) Pay(amount float64) {
	fmt.Printf("Paying %.2f using Bkash with API Key: %s\n", amount, bk.apiKey)
}

func (ng *Nagad) Pay(amount float64) {
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

func (ps *PaymentService) Checkout() {
	ps.method.Pay(10000.000)
}

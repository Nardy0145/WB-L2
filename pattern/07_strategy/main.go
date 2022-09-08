package main

const(

	)

func main() {
	product := "vehicle"
	payWay := 3
	var payment Payment
	switch payWay {
	case 1:
		payment = newCardPayment()
	case 2:
		payment = newPaypalPayment()
	case 3:
		payment = newQIWIPayment()
	}
	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	// ... implementation
	err := payment.Pay()
	if err != nil {
		return
	}
}

type Payment interface {
	Pay() error
}

type cardPayment struct {
}

func newCardPayment() Payment {
	return &cardPayment{}
}
func (p *cardPayment) Pay() error {
	// ...
	return nil
}

type paypalPayment struct {
}

func newPaypalPayment() Payment {
	return &paypalPayment{}
}
func (p *paypalPayment) Pay() error {
	// ...
	return nil
}

type qiwiPayment struct {
}

func newQIWIPayment() Payment {
	return &qiwiPayment{}
}

func (p *qiwiPayment) Pay() error {
	// ...
	return nil
}

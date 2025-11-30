package main

import (
	"fmt"
)

// The Factory Method is a creational design pattern that defines an interface for creating
// objects, while allowing subclasses (or other concrete types) to decide which object to create.

type Payment interface {
	Pay(float64) string
}

type CreateCardPay struct {
}

func (c *CreateCardPay) Pay(amt float64) string {
	return fmt.Sprintf("Credit card Payment %.2f successful", amt)
}




type RazorPay struct{}

func (r *RazorPay) Pay(amt float64) string {
	return fmt.Sprintf("RazorPay Payment %.2f successful", amt)
}




type PayPal struct{}

func (r *PayPal) Pay(amt float64) string {
	return fmt.Sprintf("RazorPay Payment %.2f successful", amt)
}



func GetPaymentFactory(payMethod string) Payment {
	switch payMethod {
	case "credit_card":
		return &CreateCardPay{}
	case "razorpay":
		return &RazorPay{}
	case "paypal":
		return &RazorPay{}
	default:
		return nil
	}
}

func main() {
	PaymentMethod := GetPaymentFactory("credit_card")
	fmt.Println(PaymentMethod.Pay(1000))

	PaymentMethod2 := GetPaymentFactory("razorpay")
	fmt.Println(PaymentMethod2.Pay(1400))

	PaymentMethod3 := GetPaymentFactory("paypal")
	fmt.Println(PaymentMethod3.Pay(100))

}

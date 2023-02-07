package main

import "fmt"

type PaymentGatewayType uint8

const (
	cash PaymentGatewayType = iota
	creditCard
	debitCard
)

type PaymentGatewayFactory struct {
	paymentGatewayMap map[PaymentGatewayType]PaymentGateway
}

func (this *PaymentGatewayFactory) getPaymentGatewayInstanceByPGType(pgType PaymentGatewayType) PaymentGateway {
	if this.paymentGatewayMap[pgType] != nil {
		return this.paymentGatewayMap[pgType]
	}
	if pgType == cash {
		this.paymentGatewayMap[pgType] = &CashPaymentGateway{}
		return this.paymentGatewayMap[pgType]
	}
	if pgType == creditCard {
		this.paymentGatewayMap[pgType] = &CreditCardPaymentGateway{}
		return this.paymentGatewayMap[pgType]
	}
	if pgType == debitCard {
		this.paymentGatewayMap[pgType] = &DebitCardPaymentGateway{}
		return this.paymentGatewayMap[pgType]
	}
	return nil
}

type PaymentGateway interface {
	pay(int)
}

type CashPaymentGateway struct {
}

func (this CashPaymentGateway) pay(price int) {
	fmt.Printf("Paying price of %d$ through cash payment\n", price)
}

type CreditCardPaymentGateway struct {
}

func (this CreditCardPaymentGateway) pay(price int) {
	fmt.Printf("Paying price of %d$ through credit card payment\n", price)
}

type DebitCardPaymentGateway struct {
}

func (this DebitCardPaymentGateway) pay(price int) {
	fmt.Printf("Paying price of %d$ through debit card payment\n", price)
}

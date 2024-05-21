package patterns

import "fmt"

/*
	Паттерн "Стратегия" определяет семейство алгоритмов,
	инкапсулирующем каждый из них и делает их взаимозаменяемыми.
	Стратегия позволяет выбирать алгоритм поведения во время выполнения,
	не изменяя классы, которые используют эти алгоритмы.

	Плюсы:
	Инкапсуляция алгоритмов
	Расширяемость
	Уменьшение дублирования

	Минусы:
	Усложнение кода
	Избыточность
*/

// ? Интерфейс стратегии оплаты

type PaymentStrategy interface {
	pay(amount float64)
}

// ? Процессор платежей

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func (p *PaymentProcessor) setStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentProcessor) processPayment(amount float64) {
	p.strategy.pay(amount)
}

// ? Оплата кредитной картой

type CreditCardPayment struct {
	cardNumber string
}

func (c *CreditCardPayment) pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card (Card Number: %s)\n", amount, c.cardNumber)
}

// ? Оплата через PayPal

type PayPalPayment struct {
	email string
}

func (p *PayPalPayment) pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal (Email: %s)\n", amount, p.email)
}

// ? Оплата через Google Pay

type GooglePayPayment struct {
	accountID string
}

func (g *GooglePayPayment) pay(amount float64) {
	fmt.Printf("Paid %.2f using Google Pay (Account ID: %s)\n", amount, g.accountID)
}

func CheckStrategy() {
	processor := &PaymentProcessor{}

	creditCardPayment := &CreditCardPayment{cardNumber: "1234-5678-9012-3456"}
	processor.setStrategy(creditCardPayment)
	processor.processPayment(100.0)

	payPalPayment := &PayPalPayment{email: "user@example.com"}
	processor.setStrategy(payPalPayment)
	processor.processPayment(200.0)

	googlePayPayment := &GooglePayPayment{accountID: "google-account-123"}
	processor.setStrategy(googlePayPayment)
	processor.processPayment(300.0)
}

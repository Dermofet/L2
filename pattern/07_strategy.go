package pattern

import "fmt"

// Интерфейс стратегии
type PaymentStrategy interface {
    Pay(amount float64)
}

// Конкретная стратегия оплаты кредитной картой
type CreditCardStrategy struct {
    cardNumber      string
    expirationDate  string
    cvv             string
}

func NewCreditCardStrategy(cardNumber, expirationDate, cvv string) *CreditCardStrategy {
    return &CreditCardStrategy{
        cardNumber:     cardNumber,
        expirationDate: expirationDate,
        cvv:            cvv,
    }
}

func (ccs *CreditCardStrategy) Pay(amount float64) {
    fmt.Printf("Оплата через кредитную карту на сумму: %.2f\n", amount)
}

// Конкретная стратегия оплаты через PayPal
type PayPalStrategy struct {
    email    string
    password string
}

func NewPayPalStrategy(email, password string) *PayPalStrategy {
    return &PayPalStrategy{
        email:    email,
        password: password,
    }
}

func (pps *PayPalStrategy) Pay(amount float64) {
    fmt.Printf("Оплата через PayPal на сумму: %.2f\n", amount)
}

// Контекст, использующий стратегию оплаты
type ShoppingCart struct {
    paymentStrategy PaymentStrategy
}

func (sc *ShoppingCart) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
    sc.paymentStrategy = paymentStrategy
}

func (sc *ShoppingCart) Checkout(amount float64) {
    if sc.paymentStrategy == nil {
        fmt.Println("Ошибка: не задана стратегия оплаты")
        return
    }
    sc.paymentStrategy.Pay(amount)
}

func RunPatternStrategy() {
    // Создание контекста (корзины покупок)
    cart := &ShoppingCart{}

    // Установка стратегии оплаты через кредитную карту
    cart.SetPaymentStrategy(NewCreditCardStrategy("1234 5678 9101 1121", "12/24", "123"))

    // Оплата с использованием текущей стратегии
    cart.Checkout(100.50)

    // Смена стратегии оплаты на PayPal
    cart.SetPaymentStrategy(NewPayPalStrategy("example@example.com", "password123"))

    // Оплата с использованием текущей стратегии
    cart.Checkout(75.25)
}

package pattern

import "fmt"

// Интерфейс продукта
type Product interface {
	GetName() string
}

// Конкретный продукт A
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "Product A"
}

// Конкретный продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "Product B"
}

// Интерфейс фабрики
type Factory interface {
	CreateProduct() Product
}

// Конкретная фабрика для создания продукта A
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// Конкретная фабрика для создания продукта B
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func RunPatternFactoryMethod() {
	// Создание фабрики для продукта A
	factories := []Factory{&ConcreteFactoryA{}, &ConcreteFactoryB{}}
	for _, factory := range factories {
		product := factory.CreateProduct()
		fmt.Println("Product created:", product.GetName())
	}
}

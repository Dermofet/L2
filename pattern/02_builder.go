package pattern

import "fmt"

// House представляет собой структуру для хранения информации о доме
type House struct {
	Foundation string
	Walls      string
	Roof       string
	Doors      int
	Windows    int
}

// HouseBuilder представляет собой интерфейс для строителей домов
type HouseBuilder interface {
	BuildFoundation()
	BuildWalls()
	BuildRoof()
	BuildDoors()
	BuildWindows()
	GetHouse() *House
}

// SimpleHouseBuilder представляет простой строитель домов
type SimpleHouseBuilder struct {
	house *House
}

// NewSimpleHouseBuilder создает новый экземпляр простого строителя домов
func NewSimpleHouseBuilder() *SimpleHouseBuilder {
	return &SimpleHouseBuilder{
		house: &House{},
	}
}

func (b *SimpleHouseBuilder) BuildFoundation() {
	b.house.Foundation = "Простой фундамент"
}

func (b *SimpleHouseBuilder) BuildWalls() {
	b.house.Walls = "Простые стены"
}

func (b *SimpleHouseBuilder) BuildRoof() {
	b.house.Roof = "Простая крыша"
}

func (b *SimpleHouseBuilder) BuildDoors() {
	b.house.Doors = 1
}

func (b *SimpleHouseBuilder) BuildWindows() {
	b.house.Windows = 2
}

func (b *SimpleHouseBuilder) GetHouse() *House {
	return b.house
}

// Director представляет собой директора строительства домов
type Director struct {
	builder HouseBuilder
}

// NewDirector создает новый экземпляр директора строительства домов
func NewDirector(builder HouseBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

// ConstructHouse конструирует дом с помощью текущего строителя
func (d *Director) ConstructHouse() {
	d.builder.BuildFoundation()
	d.builder.BuildWalls()
	d.builder.BuildRoof()
	d.builder.BuildDoors()
	d.builder.BuildWindows()
}

func RunPatternBuilder() {
	// Создаем строителя
	builder := NewSimpleHouseBuilder()

	// Создаем директора строительства
	director := NewDirector(builder)

	// Строим дом
	director.ConstructHouse()

	// Получаем готовый дом
	house := builder.GetHouse()

	// Выводим информацию о доме
	fmt.Println("Дом:")
	fmt.Printf("Фундамент: %s\n", house.Foundation)
	fmt.Printf("Стены: %s\n", house.Walls)
	fmt.Printf("Крыша: %s\n", house.Roof)
	fmt.Printf("Количество дверей: %d\n", house.Doors)
	fmt.Printf("Количество окон: %d\n", house.Windows)
}

package pattern

import "fmt"

// Element представляет собой интерфейс элемента, который может быть посещен
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA представляет собой конкретный элемент A
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// ConcreteElementB представляет собой конкретный элемент B
type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Visitor представляет собой интерфейс посетителя
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor представляет собой конкретного посетителя
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Посетитель посещает ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Посетитель посещает ConcreteElementB")
}

// ObjectStructure представляет собой объектную структуру, которую можно посещать
type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) Attach(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Detach(element Element) {
	for i, e := range os.elements {
		if e == element {
			os.elements = append(os.elements[:i], os.elements[i+1:]...)
			break
		}
	}
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range os.elements {
		element.Accept(visitor)
	}
}

func RunPatternVisitor() {
	// Создаем экземпляры элементов
	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}

	// Создаем экземпляр посетителя
	visitor := &ConcreteVisitor{}

	// Добавляем элементы в объектную структуру
	objectStructure := &ObjectStructure{}
	objectStructure.Attach(elementA)
	objectStructure.Attach(elementB)

	// Посещаем элементы с помощью посетителя
	objectStructure.Accept(visitor)
}

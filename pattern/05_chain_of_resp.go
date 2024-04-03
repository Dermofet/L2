package pattern

import "fmt"

// Request представляет собой запрос, который проходит через цепочку обязанностей
type Request struct {
	content string
}

// Handler представляет собой интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)         // Устанавливает следующий обработчик в цепочке
	HandleRequest(request *Request)  // Обрабатывает запрос
	CanHandle(request *Request) bool // Проверяет, может ли обработчик обработать запрос
}

// ConcreteHandlerA представляет конкретного обработчика A в цепочке
type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerA) CanHandle(request *Request) bool {
	// В этом примере ConcreteHandlerA обрабатывает запросы, содержащие "A"
	return request.content == "A"
}

// HandleRequest обрабатывает запрос и, если не может обработать его сам, передает следующему обработчику в цепочке
func (h *ConcreteHandlerA) HandleRequest(request *Request) {
	// Проверяем, может ли текущий обработчик обработать запрос
	if h.CanHandle(request) {
		fmt.Println("Обработчик A обрабатывает запрос:", request.content)
	} else if h.next != nil {
		// Если текущий обработчик не может обработать запрос, передаем его следующему обработчику в цепочке
		fmt.Println("Обработчик A не может обработать запрос, передаем следующему обработчику")
		h.next.HandleRequest(request)
	} else {
		// Если нет следующего обработчика, запрос остается необработанным
		fmt.Println("Ни один обработчик не может обработать запрос")
	}
}

// ConcreteHandlerB представляет конкретного обработчика B в цепочке
type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}

func (h *ConcreteHandlerB) CanHandle(request *Request) bool {
	// В этом примере ConcreteHandlerB обрабатывает запросы, содержащие "B"
	return request.content == "B"
}

// HandleRequest обрабатывает запрос и, если не может обработать его сам, передает следующему обработчику в цепочке
func (h *ConcreteHandlerB) HandleRequest(request *Request) {
	// Проверяем, может ли текущий обработчик обработать запрос
	if h.CanHandle(request) {
		fmt.Println("Обработчик B обрабатывает запрос:", request.content)
	} else if h.next != nil {
		// Если текущий обработчик не может обработать запрос, передаем его следующему обработчику в цепочке
		fmt.Println("Обработчик B не может обработать запрос, передаем следующему обработчику")
		h.next.HandleRequest(request)
	} else {
		// Если нет следующего обработчика, запрос остается необработанным
		fmt.Println("Ни один обработчик не может обработать запрос")
	}
}

func RunPatternChainOfResp() {
	// Создаем обработчики
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	// Устанавливаем следующий обработчик для handlerA
	handlerA.SetNext(handlerB)

	// Создаем запросы
	request1 := &Request{content: "A"}
	request2 := &Request{content: "B"}
	request3 := &Request{content: "C"}

	// Обрабатываем запросы
	fmt.Printf("Обработка запроса: %s\n", request1.content)
	handlerA.HandleRequest(request1)
	fmt.Printf("\nОбработка запроса: %s\n", request2.content)
	handlerA.HandleRequest(request2)
	fmt.Printf("\nОбработка запроса: %s\n", request3.content)
	handlerA.HandleRequest(request3)
}

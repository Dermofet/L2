package pattern

import "fmt"

// Command представляет собой интерфейс команды
type Command interface {
    Execute()
}

// Receiver представляет собой получателя команды
type Receiver struct{}

func (r *Receiver) Action() {
    fmt.Println("Receiver выполняет действие")
}

// ConcreteCommand представляет собой конкретную команду
type ConcreteCommand struct {
    receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
    return &ConcreteCommand{
        receiver: receiver,
    }
}

func (c *ConcreteCommand) Execute() {
    c.receiver.Action()
}

// Invoker представляет собой инициатор команды
type Invoker struct {
    command Command
}

func NewInvoker(command Command) *Invoker {
    return &Invoker{
        command: command,
    }
}

func (i *Invoker) ExecuteCommand() {
    i.command.Execute()
}

func RunPatternCommand() {
    // Создаем получателя команды
    receiver := &Receiver{}

    // Создаем конкретную команду с получателем
    command := NewConcreteCommand(receiver)

    // Создаем инициатора команды с конкретной командой
    invoker := NewInvoker(command)

    // Вызываем команду через инициатора
    invoker.ExecuteCommand()
}


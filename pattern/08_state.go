package pattern

import "fmt"

// Интерфейс состояния
type State interface {
	InsertCoin(vendingMachine *VendingMachine)   // Вносит монету
	PressButton(vendingMachine *VendingMachine)  // Нажимает кнопку
	DispenseItem(vendingMachine *VendingMachine) // Выдает товар
}

// Конкретное состояние: товар отсутствует
type SoldOutState struct{}

// Реализация метода InsertCoin для состояния SoldOutState
func (s *SoldOutState) InsertCoin(vendingMachine *VendingMachine) {
	fmt.Println("Товара нет в наличии")
}

// Реализация метода PressButton для состояния SoldOutState
func (s *SoldOutState) PressButton(vendingMachine *VendingMachine) {
	fmt.Println("Товара нет в наличии")
}

// Реализация метода DispenseItem для состояния SoldOutState
func (s *SoldOutState) DispenseItem(vendingMachine *VendingMachine) {
	fmt.Println("Товара нет в наличии")
}

// Конкретное состояние: товар доступен
type HasItemState struct{}

// Реализация метода InsertCoin для состояния HasItemState
func (s *HasItemState) InsertCoin(vendingMachine *VendingMachine) {
	fmt.Println("Монета внесена")
}

// Реализация метода PressButton для состояния HasItemState
func (s *HasItemState) PressButton(vendingMachine *VendingMachine) {
	fmt.Println("Нажата кнопка")
	vendingMachine.setState(vendingMachine.getSoldState())
}

// Реализация метода DispenseItem для состояния HasItemState
func (s *HasItemState) DispenseItem(vendingMachine *VendingMachine) {
	fmt.Println("Нельзя выдать товар до оплаты")
}

// Конкретное состояние: товар продан
type SoldState struct{}

// Реализация метода InsertCoin для состояния SoldState
func (s *SoldState) InsertCoin(vendingMachine *VendingMachine) {
	fmt.Println("Ожидайте, товар выдается")
}

// Реализация метода PressButton для состояния SoldState
func (s *SoldState) PressButton(vendingMachine *VendingMachine) {
	fmt.Println("Нельзя нажать кнопку повторно")
}

// Реализация метода DispenseItem для состояния SoldState
func (s *SoldState) DispenseItem(vendingMachine *VendingMachine) {
	vendingMachine.releaseItem()
	if vendingMachine.getCount() > 0 {
		vendingMachine.setState(vendingMachine.getHasItemState())
	} else {
		fmt.Println("Товар закончился")
		vendingMachine.setState(vendingMachine.getSoldOutState())
	}
}

// Контекст: торговый автомат
type VendingMachine struct {
	soldOutState State
	hasItemState State
	soldState    State

	state State // Текущее состояние торгового автомата
	count int   // Количество товара в автомате
}

// Создание нового торгового автомата
func NewVendingMachine(count int) *VendingMachine {
	vendingMachine := &VendingMachine{
		soldOutState: &SoldOutState{},
		hasItemState: &HasItemState{},
		soldState:    &SoldState{},
		count:        count,
	}
	if count > 0 {
		vendingMachine.state = vendingMachine.getHasItemState()
	} else {
		vendingMachine.state = vendingMachine.getSoldOutState()
	}
	return vendingMachine
}

// Получение состояния "Товар отсутствует"
func (vm *VendingMachine) getSoldOutState() State {
	return vm.soldOutState
}

// Получение состояния "Товар доступен"
func (vm *VendingMachine) getHasItemState() State {
	return vm.hasItemState
}

// Получение состояния "Товар продан"
func (vm *VendingMachine) getSoldState() State {
	return vm.soldState
}

// Установка текущего состояния торгового автомата
func (vm *VendingMachine) setState(state State) {
	vm.state = state
}

// Получение количества товара в автомате
func (vm *VendingMachine) getCount() int {
	return vm.count
}

// Выдача товара из автомата
func (vm *VendingMachine) releaseItem() {
	fmt.Println("Товар выдан")
	vm.count--
}

// Внесение монеты в торговый автомат
func (vm *VendingMachine) InsertCoin() {
	vm.state.InsertCoin(vm)
}

// Нажатие кнопки торгового автомата
func (vm *VendingMachine) PressButton() {
	vm.state.PressButton(vm)
}

// Выдача товара из торгового автомата
func (vm *VendingMachine) DispenseItem() {
	vm.state.DispenseItem(vm)
}

func RunPatternState() {
	vendingMachine := NewVendingMachine(2)

	vendingMachine.InsertCoin()
	vendingMachine.PressButton()
	vendingMachine.DispenseItem()

	vendingMachine.InsertCoin()
	vendingMachine.PressButton()
	vendingMachine.DispenseItem()

	vendingMachine.InsertCoin()
	vendingMachine.PressButton()
	vendingMachine.DispenseItem()
}

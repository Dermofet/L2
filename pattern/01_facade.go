package pattern

import (
	"fmt"
	"time"
)

// OrderFacade представляет собой фасадный класс для обработки заказов
type OrderFacade struct {
	db       *Database
	cache    *Cache
	external *ExternalService
}

// Database представляет собой объект для взаимодействия с базой данных
type Database struct{}

func (db *Database) saveOrderToDB(orderID int, items []string, address string) {
	fmt.Printf("Сохранение заказа %d в базе данных: items=%v, address=%s\n", orderID, items, address)
	// Реализация сохранения заказа в базе данных
}

// Cache представляет собой объект для работы с кешем
type Cache struct{}

func (cache *Cache) saveToCache(key string, value interface{}, expiration time.Duration) {
	fmt.Printf("Сохранение данных в кеше: ключ=%s, значение=%v, время жизни=%v\n", key, value, expiration)
	// Реализация сохранения данных в кеше
}

// ExternalService представляет собой внешний сервис для выполнения дополнительных операций
type ExternalService struct{}

func (service *ExternalService) processExternalRequest() {
	fmt.Println("Обработка внешнего запроса...")
	// Реализация взаимодействия с внешним сервисом
}

// CreateOrder создает новый заказ
func (facade *OrderFacade) CreateOrder(items []string, address string) {
	// Сохраняем заказ в базе данных
	orderID := 123 // Предположим, что у нас есть какой-то идентификатор заказа
	facade.db.saveOrderToDB(orderID, items, address)

	// Сохраняем данные в кеше
	cacheKey := "order:" + fmt.Sprint(orderID)
	facade.cache.saveToCache(cacheKey, items, time.Minute*10)

	// Взаимодействуем с внешним сервисом
	facade.external.processExternalRequest()

	fmt.Println("Заказ успешно создан!")
}

func RunPatternFacade() {
	// Создаем экземпляры базы данных, кеша и внешнего сервиса
	db := &Database{}
	cache := &Cache{}
	external := &ExternalService{}

	// Создаем экземпляр фасадного класса OrderFacade
	facade := &OrderFacade{
		db:       db,
		cache:    cache,
		external: external,
	}

	// Создаем заказ с помощью фасадного класса
	facade.CreateOrder([]string{"item1", "item2"}, "123 Main St")
}

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Парсим аргументы командной строки
	url := flag.String("url", "", "URL адрес сайта для скачивания")
	output := flag.String("output", "", "Имя файла для сохранения содержимого сайта")
	flag.Parse()

	// Проверяем наличие URL
	if *url == "" {
		fmt.Println("Необходимо указать URL адрес сайта для скачивания")
		return
	}

	// Определяем имя файла для сохранения
	fileName := *output
	if fileName == "" {
		fileName = getFileName(*url)
	}

	// Открываем файл для записи
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Ошибка при создании файла: %v\n", err)
		return
	}
	defer file.Close()

	// Отправляем GET запрос к сайту
	response, err := http.Get(*url)
	if err != nil {
		fmt.Printf("Ошибка при выполнении запроса: %v\n", err)
		return
	}
	defer response.Body.Close()

	// Копируем содержимое ответа в файл
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Printf("Ошибка при копировании данных: %v\n", err)
		return
	}

	fmt.Printf("Содержимое сайта успешно сохранено в файле: %s\n", fileName)
}

// getFileName возвращает имя файла на основе URL адреса
func getFileName(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1] + ".html"
}

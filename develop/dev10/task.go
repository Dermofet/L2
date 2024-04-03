package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Парсим аргументы командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "Таймаут на подключение к серверу")
	flag.Parse()

	// Получаем хост и порт из аргументов командной строки
	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Использование: go-telnet [--timeout=<timeout>] <host> <port>")
		return
	}
	host := args[0]
	port := args[1]

	// Подключаемся к серверу
	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Printf("Ошибка подключения к серверу: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Соединение установлено. Введите текст для отправки (Ctrl+D для выхода):")

	// Горутина для чтения данных из сокета и вывода в STDOUT
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("Соединение разорвано.")
		os.Exit(0)
	}()

	// Копируем введенные данные из STDIN в сокет
	io.Copy(conn, os.Stdin)
}

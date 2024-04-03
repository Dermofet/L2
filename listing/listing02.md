Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

выведется 2, потому что в функции test() x является частью сигнатуры функции, поэтому у отложенной функции есть доступ к x
выведется 1, потому что в функции anotherTest() отложенная функция будет работать с копией x 

```

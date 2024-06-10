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
Программа выведет:
```
2
1
```
Объяснение: defer() будет выполнена в конце выполнения функции, но до возвращения результата. 
В функции test(), x++ выполнится конкретно для возвращаемого значения, а в anotherTest() будет 
инкрементирована переменная x, которая не является напрямую возвращаемым значением.
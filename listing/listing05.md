Что выведет программа? Объяснить вывод программы.
```go
package main

type customError struct {
    msg string
}

func (e *customError) Error() string {
    return e.msg
}

func test() *customError {
    {
    // do something
    }
    return nil
}

func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```
Программа выведет:
```
error
```
Объяснение: err - имеет конкретный тип customError и не имеет значения, nil не имеет ни значения ни конкретного типа.

package main

import (
    "errors"
    "fmt"
    "log"
)

func division(x, y int) (int, error) {
    if y == 0 {
        return 0, errors.New("Cannot divide by zero!")
    }
    return x / y, nil
}

func main() {
    result, err := division(2, 2)
    if err != nil {
        fmt.Println(err)
		return
    }
    fmt.Println("The result is", result)

    result2, err := division(2, 0)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("The result is", result2)
}


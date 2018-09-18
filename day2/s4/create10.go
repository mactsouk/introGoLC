package main

import (
    "fmt"
)

func main() {
    count := 10
    fmt.Printf("Going to create %d goroutines.\n", count)

    for i := 0; i < count; i++ {
        go func(x int) {
            fmt.Printf("%d ", x)
        }(i)
    }

    fmt.Println("\nExiting...")
}


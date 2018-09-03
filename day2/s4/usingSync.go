package main

import (
    "fmt"
    "sync"
)

func main() {
    count := 10
    fmt.Printf("Going to create %d goroutines.\n", count)

    var waitGroup sync.WaitGroup
    fmt.Printf("%#v\n", waitGroup)

    for i := 0; i < count; i++ {
        waitGroup.Add(1)
        go func(x int) {
            defer waitGroup.Done()
            fmt.Printf("%d ", x)
        }(i)
    }

    fmt.Printf("%#v\n", waitGroup)
    waitGroup.Wait()
    fmt.Println("\nExiting...")
}


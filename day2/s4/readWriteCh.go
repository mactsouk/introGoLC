package main

import (
    "fmt"
    "time"
)

func writeChannel(c chan<- int, x int) {
    fmt.Println(x)
    c <- x
    time.Sleep(3 * time.Second)
    fmt.Println(x)
    close(c)
}

func main() {
    c := make(chan int)

    go writeChannel(c, 10)
    // fmt.Println("Read:", <-c)
    time.Sleep(2 * time.Second)

    value, ok := <-c
    if ok {
        fmt.Println("Channel is open!")
        fmt.Println("Value: ", value)
    } else {
        fmt.Println("Channel is closed!")
    }

    time.Sleep(4 * time.Second)
    _, ok = <-c
    if ok {
        fmt.Println("Channel is open!")
    } else {
        fmt.Println("Channel is closed!")
    }
}


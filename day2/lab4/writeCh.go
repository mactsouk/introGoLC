package main

import (
    "fmt"
    "time"
)

func wCh(c chan<- int, x int) {
    fmt.Println(x)
    c <- x
    close(c)
}

func main() {
    c := make(chan int)
    go wCh(c, 10)
    time.Sleep(2 * time.Second)
    fmt.Println("Read: ", <-c)
}


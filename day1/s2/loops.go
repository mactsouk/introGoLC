package main

import (
    "fmt"
)

func main() {
    for i := 0; i < 100; i++ {
        if i%20 == 0 {
            continue
        }
        if i == 95 {
            break
        }
        fmt.Print(i, " ")
    }

    fmt.Println()
    i := 10
    for {
        if i < 0 {
            break
        }
        fmt.Print(i, " ")
        i--
    }

    fmt.Println()
    i = 0
    anExpr := true
    for ok := true; ok; ok = anExpr {
        if i > 10 {
            anExpr = false
        }

        fmt.Print(i, " ")
        i++
    }
    fmt.Println()
    fmt.Println("Like while:")

    ok := true
    i = 0
    for ok {
        if i > 10 {
            ok = false
        }

        fmt.Print(i, " ")
        i++
    }
    fmt.Println()

    anArray := [5]int{0, 1, -1, 2, -2}
    for i, v := range anArray {
        fmt.Println("ind:", i, "v: ", v)
    }
}

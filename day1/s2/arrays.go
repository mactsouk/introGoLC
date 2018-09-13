package main

import (
    "fmt"
)

func main() {
    // myArray := []int{1, 2, 4, -4}
    myArray := [4]int{1, 2, 4, -4}
    length := len(myArray)
    for i := 0; i < length; i++ {
        fmt.Printf("%d ", myArray[i])
    }
    fmt.Printf("\n")

    otherArray := [...]int{-2, 6, 7, 8}

    for _, number := range otherArray {
        fmt.Printf("%d ", number)
    }
    fmt.Printf("\n")
    twoD := [3][3]int{
               {1, 2, 3},
               {4, 5, 6},
               {7, 8, 9}}

    twoD[1][2] = 15

    for _, number := range twoD {
        for _, other := range number {
            fmt.Printf("%d ", other)
        }
    }
    fmt.Printf("\n")
}


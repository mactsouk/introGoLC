package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

func random(min, max int) int {
    return rand.Intn(max-min) + min
}

func main() {
    MIN := 0
    MAX := 0
    TOTAL := 0

   if len(os.Args) > 3 {
        temp, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	MIN = temp

        temp, err = strconv.Atoi(os.Args[2])
        if err != nil {
                fmt.Println(err)
                os.Exit(-1)
        }
        MAX = temp

        temp, err = strconv.Atoi(os.Args[3])
        if err != nil {
                fmt.Println(err)
                os.Exit(-1)
        }
	TOTAL = temp
    } else {
        fmt.Println("Usage:", os.Args[0], "MIN MAX TOTAL")
        os.Exit(-1)
    }

    rand.Seed(time.Now().Unix())
    for i := 0; i < TOTAL; i++ {
        myrand := random(MIN, MAX)
        fmt.Print(myrand)
        fmt.Print(" ")
    }
    fmt.Println()
}


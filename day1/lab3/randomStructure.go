package main

import (
	"fmt"
	"math/rand"
	"time"
)

type XYZ struct {
	seed   int64
	first  int
	second int
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	data := []XYZ{}
	for i := 0; i < 100; i++ {
		seed := time.Now().UnixNano()
		rand.Seed(seed)
		first := random(0, 100)
		second := random(0, 100)
		temp := XYZ{seed, first, second}
		data = append(data, temp)
	}
	fmt.Println(data)
}


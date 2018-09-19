package main

import (
	"fmt"
	"math/rand"
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
	var i int64
	for i = 0; i < 100; i++ {
		rand.Seed(i)
		first := random(0, 100)
		second := random(0, 100)
		temp := XYZ{i, first, second}
		data = append(data, temp)
	}
	fmt.Println(data)
}

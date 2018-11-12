package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{}
	aMap["test"] = 1
	fmt.Println(aMap)
	aMap = nil
	fmt.Println(aMap)
	aMap["test"] = 1
}

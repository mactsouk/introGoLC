package aPackage

import (
	"fmt"
)

func A() {
	fmt.Println("This is function A! - CHANGED")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}

const MyConstant = 123
const privateConstant = 21



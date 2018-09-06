package main

import (
	"fmt"
)

func main() {
	var temperature float64
	var t int

	for {
		fmt.Println("Give me the temperature: ")
		fmt.Scanf("%f", &temperature)
		fmt.Println("0 for Celsius - 1 for Fahrenheit - anything else to exit")
		fmt.Scanf("%d", &t)
		if t == 0 {
			fmt.Println("Converting Celsius to Fahrenheit.")
			f := 1.8*temperature + 32
			fmt.Printf("%.2f C is %.3f F.\n", temperature, f)
		} else if t == 1 {
			fmt.Println("Converting Fahrenheit to Celsius.")
			c := (temperature - 32) * 5 / 9
			fmt.Printf("%.2f F is %.3f C.\n", temperature, c)
		} else {
			return
		}
	}
}

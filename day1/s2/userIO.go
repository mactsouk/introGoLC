package main

import "fmt"

func main() {
    fmt.Print("Give me a number: ")
    var aNumber int
    fmt.Scanln(&aNumber)
    fmt.Println("Your number is:", aNumber)

    // Read a string
    fmt.Print("Enter your Name: ")
    var name string
    fmt.Scanln(&name)
    fmt.Println("Your name is:", name)
    fmt.Print("Your name is: ", name, "\n")
    fmt.Printf("Your name is: %s\n", name)

    // Print its type
    fmt.Printf("The type of variable name is %T\n", name)
    // Define the decimal points of the output
    fmt.Printf("%.4f\n", 123.1231231)
    fmt.Println(123.1231231)
}


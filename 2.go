package main

import "fmt"

func main() {
    var c, f float64
    fmt.Print("Digite a temperatura em Celsius: ")
    fmt.Scanln(&c)
    f = (c * 9 / 5) + 32
    fmt.Println("A temperatura em Fahrenheit é:", f)
}

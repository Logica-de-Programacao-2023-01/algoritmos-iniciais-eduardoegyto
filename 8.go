package main

import "fmt"

func triangulo(base, altura float64) float64 {
    return (base * altura) / 2
}

func main() {
    var base, altura float64
    fmt.Print("Digite a base do triângulo: ")
    fmt.Scanln(&base)
    fmt.Print("Digite a altura do triângulo: ")
    fmt.Scanln(&altura)
    area := triangulo(base, altura)
    fmt.Println("A área do triângulo é:", area)
}

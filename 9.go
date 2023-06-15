package main

import (
    "fmt"
    "math"
)

func juros(valor, taxa, tempo float64) float64 {
    resultado := valor * math.Pow((1 + taxa/100), tempo)
    return resultado
}

func main() {
    var valor, taxa, tempo float64
    fmt.Print("Digite o valor principal: ")
    fmt.Scanln(&valor)
    fmt.Print("Digite a taxa de juros: ")
    fmt.Scanln(&taxa)
    fmt.Print("Digite o período de tempo (em anos): ")
    fmt.Scanln(&tempo)
    resultado := juros(valor, taxa, tempo)
    fmt.Println("O montante após o período de tempo é:", resultado)
}

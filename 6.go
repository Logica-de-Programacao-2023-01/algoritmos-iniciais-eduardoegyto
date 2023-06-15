package main

import "fmt"

func potencia(base, exponente int) int {
    resultado := 1
    for i := 0; i < exponente; i++ {
        resultado *= base
    }
    return resultado
}

func main() {
    var base, exponente int
    fmt.Print("Digite a base: ")
    fmt.Scanln(&base)
    fmt.Print("Digite o expoente: ")
    fmt.Scanln(&exponente)
    resultado := potencia(base, exponente)
    fmt.Println("O resultado da potência é:", resultado)
}

package main

import (
    "fmt"
    "math"
)

func primo(x int) bool {
    if x <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
        if x%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    var x int
    fmt.Print("Digite um número: ")
    fmt.Scanln(&x)

    if primo(x) {
        fmt.Println("O número", x, "é primo.")
    } else {
        fmt.Println("O número", x, "não é primo.")
    }
}

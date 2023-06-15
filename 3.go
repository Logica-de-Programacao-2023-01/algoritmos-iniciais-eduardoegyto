package main

import "fmt"

func main() {
    var x int
    fmt.Print("Digite a quantidade de números: ")
    fmt.Scanln(&x)

    var numeros []float64
    for i := 0; i < x; i++ {
        var num float64
        fmt.Printf("Digite o número %d: ", i+1)
        fmt.Scanln(&num)
        numeros = append(numeros, num)
    }

    soma := 0.0
    for _, num := range numeros {
        soma += num
    }
    media := soma / float64(x)
    fmt.Println("A média dos números é:", media)
}

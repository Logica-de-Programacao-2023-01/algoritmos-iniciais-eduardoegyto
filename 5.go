package main

import "fmt"

func busca(arr []int, key int) int {
    for i, val := range arr {
        if val == key {
            return i
        }
    }
    return -1
}

func main() {
    x := []int{2, 4, 6, 8, 10}
    fmt.Println("Lista de números:", x)

    var y int
    fmt.Print("Digite o número a ser pesquisado: ")
    fmt.Scanln(&y)

    z := busca(x,y)
    if z != -1 {
        fmt.Println("O número", y, "foi encontrado no índice", z)
    } else {
        fmt.Println("O número", y, "não foi encontrado na lista")
    }
}

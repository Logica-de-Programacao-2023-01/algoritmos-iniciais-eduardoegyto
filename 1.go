package main

import "fmt"

func main() {
    var x, y int
    fmt.Print("Digite o primeiro número: ")
    fmt.Scanln(&x)
    fmt.Print("Digite o segundo número: ")
    fmt.Scanln(&y)
    soma := x + y
    fmt.Println("A soma dos números é:", soma)
}

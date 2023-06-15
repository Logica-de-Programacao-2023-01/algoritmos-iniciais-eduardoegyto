package main

import "fmt"

func main() {
    var x int
    fmt.Print("De qual número você quer saber o fatorial? ")
    fmt.Scanln(&x)

    fatorial := 1
    for i := 1; i <= x; i++ {
        fatorial *= i
    }
    fmt.Println("O fatorial de", x, "é:", fatorial)
}

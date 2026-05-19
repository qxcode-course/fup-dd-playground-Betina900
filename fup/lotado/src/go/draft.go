package main

import "fmt"

func main() {
    var capacidade int
    fmt.Scan(&capacidade)

    passageiros := 0

    for {
        var movimento int
        fmt.Scan(&movimento)

        passageiros += movimento

        if passageiros == 0 {
            fmt.Println("vazio")

        } else if passageiros >= 2*capacidade {
            fmt.Println("hora de partir")
            break

        } else if passageiros >= capacidade {
            fmt.Println("lotado")

        } else {
            fmt.Println("ainda cabe")
        }
    }
}
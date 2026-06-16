package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    for linha := 0; linha < n; linha++ {

        // pontos da esquerda
        for i := 0; i < n-linha-1; i++ {
            fmt.Print(".")
        }

        // números
        for i := 0; i < linha+1; i++ {
            if i > 0 {
                fmt.Print(".")
            }
            fmt.Print(n)
        }

        // pontos da direita
        for i := 0; i < n-linha-1; i++ {
            fmt.Print(".")
        }

        fmt.Println()
    }
}
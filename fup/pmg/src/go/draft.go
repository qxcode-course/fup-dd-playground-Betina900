package main
import "fmt"
func main() {
    var soldados int
    fmt.Scan(&soldados)

    alturas := make([]float64, soldados)
    soma := 0.0
    for i := 0; i < len(alturas); i++{
        fmt.Scan(&alturas[i])
        soma += alturas[i]
    }

    media := soma / float64(soldados)
    fmt.Printf("%.2f\n", media)

    for i := 0; i < len(alturas); i++{
        var letra string

        if alturas[i] == media{
            letra = "M"
        } else if alturas[i] < media{
            letra = "P"
        } else{
            letra = "G"
        }

        if i == len(alturas) - 1{
            fmt.Print(letra)
        } else{
            fmt.Print(letra, " ")
        }
    }
    fmt.Println()
}
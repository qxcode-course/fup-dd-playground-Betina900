package main
import "fmt"
func main() {
    var P, S, E int
    fmt.Scan(&P, &S, &E)

    posicao := 0

    for{
        inicio := posicao
        posicao += S

        // saiu do poço
        if posicao >= P {
            fmt.Println(inicio, "saiu")
            break
        }

        // imprime salto
        fmt.Println(inicio, posicao)

        // escorrega
        posicao -= E
    }
}

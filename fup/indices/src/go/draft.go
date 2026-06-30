package main
import (
    "fmt"
    "sort"
)

type Numero struct{
    valor int
    indice int
}
func main() {
    var n int
    fmt.Scan(&n)

    vetor := make([]Numero, n)
    for i := 0; i < n; i++{
        fmt.Scan(&vetor[i].valor)
        vetor[i].indice = i
    }

    sort.Slice(vetor, func (i, j int) bool{
        return vetor[i].valor < vetor[j].valor
    })

    fmt.Print("[ ")
    for _, num := range vetor{
            fmt.Print(num.indice, " ")
    }
    fmt.Println("]")
}
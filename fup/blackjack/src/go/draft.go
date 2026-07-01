package main
import "fmt"
func main() {
    var q int
    fmt.Scan(&q)

    vetor := make([]int, q)

    soma := 0
    ases := 0
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])

        if vetor[i] == 1{
            soma += 11
            ases++
        } else if vetor[i] == 11 || vetor[i] == 12 || vetor[i] == 13{
            soma += 10
        } else{
            soma += vetor[i]
        }
    }

    for soma > 21 && ases > 0{
        soma -= 10
        ases--

    }

    fmt.Println(soma)
}
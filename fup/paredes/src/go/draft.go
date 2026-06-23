package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    vetor := make([]int, n)
    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])
    }

    maior := vetor[0]
    count := 1
    
    for i := 1; i < len(vetor); i++ {
        if vetor[i] > maior{
            count++
            maior = vetor[i]
        }
    }

    fmt.Println(count)
}
package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    maior := 1
    count := 1

    vetor := make([]int, n)
    for i := 0; i <= n; i++{
        fmt.Scan(&vetor[i])

        if vetor[i] > maior{
            count++
            maior = vetor[i]
        }
    }

    fmt.Println(count)
}
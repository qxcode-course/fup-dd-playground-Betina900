package main
import "fmt"
func main() {
    var t int
    var l1, l2 int
    fmt.Scan(&t, &l1, &l2)

    vetor := make([]int, t)

    count := 0

    for i := 0; i < len(vetor); i++{
        fmt.Scan(&vetor[i])

        if (vetor[i] >= l1 && vetor[i] <= l2){
            count += 1
        }
    }

    fmt.Println(count)
}

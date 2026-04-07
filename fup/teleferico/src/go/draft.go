package main
import (
    "fmt"
    "math"
)
func main() {
    var cap_cabine, num_alunos float64
    fmt.Scan(&cap_cabine, &num_alunos)

    quantViagens:=num_alunos/(cap_cabine-1)
    fmt.Println(math.Ceil(quantViagens))
}

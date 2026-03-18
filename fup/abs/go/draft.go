package main
import (
    "fmt"
    "math"
)
func main() {
    var valor1, valor2 float64
    fmt.Scan(&valor1, &valor2)

    diferenca := math.Abs(valor1 - valor2)
    fmt.Println(diferenca)
}

package main
import (
    "fmt"
    "math"
)
func main() {
    var x1, y1, x2, y2 float64
    fmt.Scan(&x1, &y1, &x2, &y2)

    var dab = ((x2-x1)*(x2-x1))+((y2-y1)*(y2-y1))
    var raiz = math.Sqrt(dab)
    fmt.Printf("%.2f\n", raiz)
}

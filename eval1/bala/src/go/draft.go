package main
import (
    "fmt"
    "math"
)
func main() {
    var coo_x1, coo_y1, coo_x2, coo_y2, resultado float64
    fmt.Scan(&coo_x1, &coo_y1, &coo_x2, &coo_y2)

    resultado=math.Sqrt(((coo_x2-coo_x1)*(coo_x2-coo_x1))+((coo_y2-coo_y1)*(coo_y2-coo_y1)))
    fmt.Printf("%.02f\n", resultado)
}

package main
import (
	"fmt"
	"math"
)
func main() {
    var a, b, c float64
	fmt.Scan(&a, &b, &c)

	var p = (a+b+c)/2
	var area = ((p-a)*(p-b)*(p-c))*p
	var raiz = math.Sqrt(area)
	fmt.Printf("%.2f\n", raiz)
}

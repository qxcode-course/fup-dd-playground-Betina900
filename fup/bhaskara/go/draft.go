package main
import (
    "fmt"
    "math"
)
func main() {
    var valorA, valorB, valorC float64
    fmt.Scan(&valorA, &valorB, &valorC)
    
    delta := (math.Pow(valorB, 2))-(4*valorA*valorC)
    x1 := ((-valorB)+(math.Sqrt(delta)))/(2*valorA)
    x2 := ((-valorB)-(math.Sqrt(delta)))/(2*valorA)
    if delta>0{
        fmt.Printf("%.02f\n", x1)
        fmt.Printf("%.02f\n", x2)
    } else if delta==0{
        fmt.Printf("%.02f\n", x1)
    } else{
        fmt.Println("nao ha raiz real")
    }
}

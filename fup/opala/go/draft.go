package main
import "fmt"
func main() {
    var km, min, litro float64
    fmt.Scan(&km, &min, &litro)
    hora:= min/60
    dist:= km*hora
    desem:= dist/litro
    fmt.Printf("%.2f\n", float64(desem))
}

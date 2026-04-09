package main
import "fmt"
func main() {
    var salario  float64
    fmt.Scan(&salario)
    if salario<=1000{
        salario = salario+(salario*0.2)
        fmt.Printf("%.02f\n", salario)
    } else if salario<=1500{
        salario = salario+(salario*0.15)
        fmt.Printf("%.02f\n", salario)
    } else if salario<=2000{
        salario = salario+(salario*0.1)
        fmt.Printf("%.02f\n", salario)
    } else{
        salario = salario+(salario*0.05)
        fmt.Printf("%.02f\n", salario)
    }
}

package main
import "fmt"
func main() {
    var numero1, numero2 int
    fmt.Scan(&numero1, &numero2)

    divisao_inteira := numero1/numero2
    fmt.Println(divisao_inteira)
    resto_divisao := numero1%numero2
    fmt.Println(resto_divisao)
    resultado := float64(numero1) / float64(numero2)
    fmt.Printf("%.2f\n", resultado)
}

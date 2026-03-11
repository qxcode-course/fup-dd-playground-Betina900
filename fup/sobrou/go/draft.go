package main
import "fmt"
func main() {
    var produt1, produt2, produt3 float64
    var preco1, preco2, preco3 float64
    var dinheiro float64
    fmt.Scan(&produt1, &produt2, &produt3, &preco1, &preco2, &preco3, &dinheiro)

    var valor = ((produt1*preco1)+(produt2*preco2)+(produt3*preco3))
    fmt.Printf("%.2f\n", dinheiro-valor)
}

package main
import "fmt"
func main() {
    var numero1, numero2 int
    fmt.Scan(&numero1, &numero2)

    //soma
    soma:=numero1+numero2
    fmt.Println(soma)
    //diferença
    diferenca:=numero1-numero2
    fmt.Println(diferenca)
    //multipliação
    mult:=numero1*numero2
    fmt.Println(mult)
    //divisão
    divisao:=float64(numero1)/float64(numero2)
    fmt.Printf("%.2f\n", divisao)
    //resto
    resto:=numero1%numero2
    fmt.Println(resto)
}

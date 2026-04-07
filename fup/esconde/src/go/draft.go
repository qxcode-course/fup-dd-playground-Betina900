package main
import "fmt"
func imprimir_impares(num_inteiro int) {
    for i:=1; i<=num_inteiro; i++{
        if i%2==1{
            fmt.Println(i)
        }
    }
}

func imprimir_pares(num_inteiro int) {
    for i:=num_inteiro; i>=0; i--{
        if i%2==0{
            fmt.Println(i)
        }
    }
}

func main() {
    var num_inteiro int
    fmt.Scan(&num_inteiro)
    imprimir_impares(num_inteiro)
    imprimir_pares(num_inteiro)
}

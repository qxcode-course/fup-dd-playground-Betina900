package main
import "fmt"
func main() {
    var valor1, valor2, valor3 int
    fmt.Scan(&valor1, &valor2, &valor3)

    if valor1+valor2<=valor3 || valor1+valor3<=valor2 || valor2+valor3<=valor1{
        fmt.Println("False")
    } else{
        fmt.Println("True")
    }
}

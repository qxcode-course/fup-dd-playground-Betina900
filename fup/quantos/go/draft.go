package main
import "fmt"
func main() {
    var valor1 int
    var valor2 int
    var valor3 int
    fmt.Scan(&valor1, &valor2, &valor3)

    if valor1==valor2 && valor1==valor3 {
        fmt.Println(3)
    } else if valor1==valor2 && valor1!=valor3{
        fmt.Println(2)
    } else if valor1!=valor2 && valor1==valor3 {
        fmt.Println(2)
    } else if valor2==valor3 && valor2==valor1 {
        fmt.Println(3)
    } else if valor2==valor3 && valor2!=valor1 {
        fmt.Println(2)
    } else {
        fmt.Println(0)
    }
}

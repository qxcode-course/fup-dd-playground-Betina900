package main
import "fmt"
func main() {
    var valorA, valorB int
    fmt.Scan(&valorA, &valorB)

    soma_valor := 0

    if valorB<valorA{
        fmt.Println("invalido")
        return
    }

    for i:=valorB; i>=valorA; i--{
        if i%2==0 && i%3==0{
            soma_valor += i
        }
    }
    fmt.Println(soma_valor)
}

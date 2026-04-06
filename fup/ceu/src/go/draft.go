package main
import "fmt"

func imprimir_sequencia(num_pedra int) {
    fmt.Print("[ ")
    for i:=0; i<=10; i++{
        if num_pedra==i{
            continue
        }
        if i==10{
            fmt.Print("ceu ")
            continue
        }

            fmt.Print(i, " ")
    }
    fmt.Println("]")
}
func main() {
    var num_pedra int
    fmt.Scan(&num_pedra)

    imprimir_sequencia(num_pedra)
}

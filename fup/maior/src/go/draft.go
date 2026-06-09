package main
import "fmt"
func main() {
    var chute float64
    var escolha string
    var real float64
    fmt.Scan(&chute, &escolha, &real)

    if escolha == "m"{
        if chute == real || chute < real{
            fmt.Println("primeiro")
        } else{
            fmt.Println("segundo")
        }
    } else {
        if chute == real || chute > real{
            fmt.Println("primeiro")
        } else{
            fmt.Println("segundo")
        }
    }
}
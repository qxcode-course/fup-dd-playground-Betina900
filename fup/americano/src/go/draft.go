package main
import "fmt"
func main() {
    var num1 int
    var num2 int
    var num3 int
    var num4 int

    fmt.Scan(&num1, &num2, &num3, &num4)

    soma := num1 + num2 + num3 + num4

    if soma == 0{
        fmt.Println("nenhum")
    }

    //vencedor := " "
}
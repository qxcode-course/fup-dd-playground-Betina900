package main
import "fmt"
func main() {
    var num1, num2 int
    var op string
    fmt.Scan(&num1, &num2, &op)

    if op=="+" {
        fmt.Println(num1+num2)
    } else if op =="-" {
        fmt.Println(num1-num2)
    } else if op=="/"{
        fmt.Println(num1/num2)
    } else {
        fmt.Println(num1*num2)
    }
}

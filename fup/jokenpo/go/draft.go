package main
import "fmt"
func main() {
    var op1, op2 string
    fmt.Scan(&op1, &op2)

    if op1=="R" && op2=="S" || op1=="S" && op2=="P" || op1=="P" && op2=="R" {
        fmt.Println("jog1")
    } else if op2=="R" && op1=="S" || op2=="S" && op1=="P" || op2=="P" && op1=="R"{
        fmt.Println("jog2")
    } else {
        fmt.Println("empate")
    }
}

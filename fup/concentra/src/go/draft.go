package main
import "fmt"
func main() {
    var numA, numB int
    fmt.Scan(&numA, &numB)
    
    A := numA
    B := numB
    fmt.Print("[ ")
    for i:=numA; i<=numB; i++{
        fmt.Print(A, " ")
        A += 1
        fmt.Print(B, " ")
        B -= 1
    }
    fmt.Print("]\n")
}

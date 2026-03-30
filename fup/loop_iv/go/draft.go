package main
import "fmt"
func main() {
    var numA, numB int
    fmt.Scan(&numA, &numB)

    fmt.Print("[ ")
    for i:=numA; i>numB; i-- {
        fmt.Print(i, " ")
    }
    for i:=numA; i<numB; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println("]")
}

package main
import "fmt"
func main() {
    var numA, numB int
    fmt.Scan(&numA, &numB)

    fmt.Print("[ ")
    for i:=numA; i<numB; i++ {
        if i % 2 == 0 {
            continue
        }
        fmt.Print(i, " ")
    }
    fmt.Println("]")
}

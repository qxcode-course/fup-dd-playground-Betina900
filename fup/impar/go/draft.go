package main
import "fmt"
func main() {
    var P, D1, D2 int
    fmt.Scan(&P, &D1, &D2)

    conta := (D1+D2)%2
    if P==0 && conta==0 {
        fmt.Println(0)
    } else if P==0 && conta==1 {
        fmt.Println(1)
    } else if P==1 && conta==1 {
        fmt.Println(0)
    } else {
        fmt.Println(1)
    }
}

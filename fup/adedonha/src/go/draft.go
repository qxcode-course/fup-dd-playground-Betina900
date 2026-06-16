package main
import "fmt"
func main() {
    var l int
    fmt.Scan(&l)

    if l == 0 {
        fmt.Println("joguem de novo")
    } else if l%26 == 0 {
        fmt.Println("z")
    } else {
        fmt.Println(string('a' + (l%26) - 1))
    }
}
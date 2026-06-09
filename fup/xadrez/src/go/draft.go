package main
import "fmt"
func main() {
    var linha int
    var coluna int
    fmt.Scan(&linha, &coluna)

    if ((linha + coluna) % 2 == 0){
        fmt.Println("1")
    } else {
        fmt.Println("0")
    }
}
package main
import "fmt"
func main() {
    var letra string
    fmt.Scan(&letra)

    c := letra[0]

    if c >= 'a' && c <= 'z'{
        fmt.Println(string(c - 32))
    } else if c >= 'A' && c <= 'Z'{
        fmt.Println(string(c + 32))
    } else{
        fmt.Println(string(c))
    }
}
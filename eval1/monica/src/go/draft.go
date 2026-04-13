package main
import "fmt"
func main() {
    var idade_monica, idade_a, idade_b, idade_c int
    fmt.Scan(&idade_monica, &idade_a, &idade_b)

    idade_c=idade_monica-(idade_a+idade_b)
    if idade_a>idade_b && idade_a>idade_c{
        fmt.Println(idade_a)
    } else if idade_b>idade_a && idade_b>idade_c{
        fmt.Println(idade_b)
    } else{
        fmt.Println(idade_c)
    }
}

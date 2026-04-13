package main
import "fmt"
func main() {
    var num_a, num_b, soma int
    fmt.Scan(&num_a, &num_b)
    if num_a>num_b{
        fmt.Println("invalido")
    } else{
        for i:=num_a; i<=num_b; i++{
            if i%2==0{
                soma=soma+i
            }
        }
        fmt.Println(soma)
    }
}

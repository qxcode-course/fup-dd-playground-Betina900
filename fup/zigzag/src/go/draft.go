package main
import "fmt"
func sequencia_jogo(num_inicio, num_fim int) {
    for i:=num_inicio; i<=num_fim; i++{
        if i%3==0 && i%5==0{
            fmt.Println("zigzag")
            continue
        }
        if i%3==0{
            fmt.Println("zig")
            continue
        }
        if i%5==0{
            fmt.Println("zag")
            continue
        }

        fmt.Println(i)
    }
}
func main() {
    var num_inicio, num_fim int
    fmt.Scan(&num_inicio, &num_fim)
    sequencia_jogo(num_inicio, num_fim)
}

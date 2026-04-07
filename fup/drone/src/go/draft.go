package main
import "fmt"
func main() {
    var a, b, c int
    var at_janela, lar_janela int
    fmt.Scan(&a, &b, &c, &at_janela, &lar_janela)

    areaJ:=lar_janela*at_janela
    if a*b<=areaJ{
        fmt.Println("S")
    } else if a*c<=areaJ{
        fmt.Println("S")
    } else if b*c<=areaJ{
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}

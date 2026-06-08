package main
import "fmt"
func main() {
    var h1, b1, h2, b2 int
    fmt.Scan(&h1,&b1, &h2, &b2)

    //o melhor local = a maior área
    //os terrenos são retangulo
    //área do retangulo é: A = b(largura) x h(comprimento)
    area1 := b1 * h1
    area2 := b2 * h2

    if area1 > area2{
        fmt.Println(area1)
    } else {
        fmt.Println(area2)
    }
}
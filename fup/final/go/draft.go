package main
import "fmt"
func main() {
    var nota1, nota2, nota_final int
    fmt.Scan(&nota1, &nota2, &nota_final)

    media := (nota1+nota2)/2
    media_final := (media+nota_final)/2

    if media>=7{
        fmt.Println("aprovado")
    } else if media<4{
        fmt.Println("reprovado")
    } else if media>4 && media<7 && media_final>=5{
        fmt.Println("aprovado na final")
    } else{
        fmt.Println("reprovado na final")
    }
}

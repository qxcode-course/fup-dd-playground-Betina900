package main
import "fmt"
func main() {
    var n1 int
    var n2 int
    var n3 int
    var n4 int
    fmt.Scan(&n1, &n2, &n3, &n4)

    /*É uma questão com pegadinha. É preciso ignorar que o jogador 1 começa
    contando no 0, porque se a soma dos dodos de cada jogador der 01 quer dizer
    que o jog1 ganhou. Por isso soma dividido para 4 (porque são quatro
    jogadores) tiver resto 1 significa que soma é igual a 1 e jog 1 ganhou.
    Resto 2 para jog2, resto 3 para jog3 e resto 4 para jog4*/

    soma := n1 + n2 + n3 + n4
    
    if soma == 0{
        fmt.Println("nenhum")
    } else if soma % 4 == 1 {
        fmt.Println("jog1")
    } else if soma % 4 == 2 {
        fmt.Println("jog2")
    } else if soma % 4 == 3 {
        fmt.Println("jog3")
    } else {
        fmt.Println("jog4")
    }

}
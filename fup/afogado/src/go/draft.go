package main
import "fmt"
func main() {
    var p int
    var e int
    fmt.Scan(&p, &e)

    /*Essa questão ensina uma técnica muito comum: For externo (Testa
    possibilidades, sem condição de parada); For interno (Simula o que
    acontece e retorna quando encontra)*/
    for i := 1; ; i++{
        posicao := 0
        salto := i

        for{
            if salto <= 0 {
                break
            }

            posicao += salto

            if posicao >= p{
                fmt.Println(i)
                return
            }

            posicao -= e

            if posicao < 0 {
                break
            }

            salto -= 10
        }
    }

}
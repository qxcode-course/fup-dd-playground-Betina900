package main
import "fmt"
func main() {
    var limite int
    var qntd int
    fmt.Scan(&limite, &qntd)

    jogadas := make([]int, qntd)

    ganhador := -1
    menorDist := limite + 1
    perdedor := 0
    maiorDist := -1

    for i := 0; i < len(jogadas); i++{
        fmt.Scan(&jogadas[i])

        distancia := jogadas[i]
        if distancia < 0{
            distancia = -distancia
        }

        if distancia <= limite{
            if distancia <= menorDist{
                menorDist = distancia
                ganhador = i
            }
        }

        if distancia >= maiorDist{
            maiorDist = distancia
            perdedor = i
        }
    }

    if ganhador == -1{
        fmt.Println("nenhum")
    } else{
        fmt.Println(ganhador)
    }

    fmt.Println(perdedor)
}
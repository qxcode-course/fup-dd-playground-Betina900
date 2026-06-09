package main
import "fmt"
func main() {
    var n_posicoes, disco, aviao int
    fmt.Scan(&n_posicoes, &disco, &aviao)
    
    distancia := (disco-aviao+n_posicoes)%n_posicoes
    fmt.Println(distancia)
}

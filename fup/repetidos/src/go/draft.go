package main
import "fmt"
func main() {
    var num_p, tamanhoVetor int
    fmt.Scan(&num_p, &tamanhoVetor)
    elementos := make([]int, tamanhoVetor)
    for i:=0; i<tamanhoVetor; i++{
        fmt.Scan(&elementos[i])
    }
    contador := 0
    for _, valor:= range elementos{
        if valor==num_p{
            contador+=1
        }
    }

    fmt.Println(contador)
}

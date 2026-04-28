package main
import (
    "fmt"
    "strconv"
)
// parâmetros: (o vetor onde vamos procurar, o números que queremos enconterar)
// retorno: "boll" retorna verdadeiro ou falso
func contains(num []int, valor int) bool{
    for _, elem := range num{
        if elem == valor{
            return true
        }
    }
    return false
}
// o nome da função (nome do vetor que ela recebe) (retorna dois valores)
// o objetivo da função é separar as figurinhas entre "album" (não repetidas)
// e "rept" (repetida)
func separar_figurinhas(montante []int) ([]int, []int) {
    // ciração de dois vetores vazios
    album := make([]int, 0, len(montante))
    repet := make([]int, 0, len(montante))
    for _, fig := range montante{
        if !contains(album, fig){
            album = append(album, fig)
        } else{
            repet = append(repet, fig)
        }
    }
    return album, repet
}

func main() {
    var montante []int = make ([]int, 0, 1)
    fmt.Println(montante, len(montante), cap(montante))

    // pega o montante e adiciona todos os números no final, depois devolve o resultado atualizado
    montante = append(montante, 7, 3, 2, 1, 9, 1, 2, 3, 4, 5, 4, 3, 2, 1, 2, 5, 7)
    // retorna dois valores: número convertido e um erro (caso dê problema)
    num, err := strconv.Atoi("32432")
    if err == nil{
        fmt.Println(num)
    } else{
        fmt.Println(err)
    }
}


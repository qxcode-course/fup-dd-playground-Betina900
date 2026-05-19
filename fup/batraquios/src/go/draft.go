package main
import "fmt"
func main() {
    //primeiro vetor
    var N1 int
    fmt.Scan(&N1)
    V1 := make([]int, N1)
    for i := 0; i < N1; i++{
        fmt.Scan(&V1[i])
    }
    //segundo vetor
    var N2 int
    fmt.Scan(&N2)
    V2 := make([]int, N2)
    for i := 0; i < N2; i++{
        fmt.Scan(&V2[i])
    }

    contido := true
    //percorre primeiro vetor
    for _, num1 := range V1{
        encontrou := false

        //procura se os elementos do primeirto vetor estão no segundo
        for _, num2 := range V2{
            if num1 == num2{
                encontrou = true
                break
            }
        }

        // se não encontrou, já pode parar
        if !encontrou {
            contido = false
            break
        }
    }

    if contido{
        fmt.Println("sim")
    } else{
        fmt.Println("nao")
    }
}

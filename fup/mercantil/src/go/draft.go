package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    valores := make([]float64, n)
    chutes := make([]float64, n)
    escolhas := make([]string, n)

    for i := 0; i < n; i++{
        fmt.Scan(&valores[i])
    }
    for i := 0; i < n; i++{
        fmt.Scan(&chutes[i])
    }
    for i := 0; i < n; i++{
        fmt.Scan(&escolhas[i])
    }

    primeiro := 0
    segundo := 0

    for i := 0; i < n; i++{
        if valores[i] == chutes[i]{
            primeiro++
        } else if escolhas[i] == "M"{
            if valores[i] > chutes[i]{
                segundo++
            } else{
                primeiro++
            }
        } else{
            if valores[i] < chutes[i]{
                segundo++
            } else{
                primeiro++
            }
        }
    }

    if primeiro > segundo{
        fmt.Println("primeiro")
    } else if segundo > primeiro{
        fmt.Println("segundo")
    } else{
        fmt.Println("empate")
    }
}
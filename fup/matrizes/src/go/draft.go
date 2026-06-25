package main
import "fmt"
func main() {
    var l, c int
    fmt.Scan(&l, &c)

    matriz := make([][]int, l)
    matrize := make([][]int, l)
    soma := make ([][]int, l)

    for i:= 0; i < l; i++{
        matriz[i] = make([]int, c)
        matrize[i] = make([]int, c)
        soma[i] = make([]int, c)
    }

    for i := 0; i < l; i++{
        for j := 0; j < c; j++{
            fmt.Scan(&matriz[i][j])
        }
    }

    for i := 0; i < l; i++{
        for j := 0; j < c; j++{
            fmt.Scan(&matrize[i][j])
        }
    }

    for i := 0; i < l; i++{
        fmt.Print("[ ")
        for j := 0; j < c; j++{
            soma[i][j] = matriz[i][j] + matrize[i][j]
            fmt.Print(soma[i][j], " ")
        }

        fmt.Print("]")
        fmt.Println()
    }
}
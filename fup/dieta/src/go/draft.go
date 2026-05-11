package main
import "fmt"
func main() {
    var dias int
    var soma float64
    fmt.Scan(&dias)

    calorias := make([]float64, dias) //exige de 'dias' (no caso) seja int

    for i:=0; i<dias; i++{
        fmt.Scan(&calorias[i])
    }

    for i:=0; i<dias; i++{
        soma += calorias[i]
    }
    media_calorias := soma/float64(dias)

    fmt.Printf("%.1f\n", media_calorias)
}

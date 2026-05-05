package main
import "fmt"
func main() {
    var dias float64
    fmt.Scan(&dias)
    calorias := make([]float64, dias)

    for i:=0; i<dias; i++{
        fmt.Scan(&calorias[i])
    }

    for i:=0; i<=dias; i++{
        calorias += calorias
    }
    media_calorias := calorias/dias
    fmt.Printf("%.1f", media_calorias)
}

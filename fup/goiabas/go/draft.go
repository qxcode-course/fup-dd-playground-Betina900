package main
import (
    "fmt"
    "math"
)
func main() {
    var capacidade, banana, goiaba, manga int
    fmt.Scan(&capacidade, &banana, &goiaba, &manga)

    total := banana+goiaba+manga
    tempo :=float64(total)/float64(capacidade)
    if capacidade>=total{
        fmt.Println(1)
    } else{
        fmt.Println(math.Ceil(tempo))
    }
}

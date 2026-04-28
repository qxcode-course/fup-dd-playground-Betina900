package main
import "fmt"
func main() {
    var quant_numeros int
    fmt.Scan(&quant_numeros)
    arr := make([]int, quant_numeros)
    
    for i := 0; i<quant_numeros; i++{
        fmt.Scan(&arr[i])
    }

    fmt.Print("[")
    for _, nums := range arr{
        fmt.Print(" ", nums)
    }
    fmt.Println(" ]")
}

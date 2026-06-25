package main
import "fmt"
func main() {
    var matriz[5][5]int

    for i := 0; i < 5; i++{
        for j := 0; j <5; j++{
            fmt.Scan(&matriz[i][j])
        }
    }

    dp := 0
    for i := 0; i < 5; i++{
        for j := 0; j < 5; j++{
            if i == j{
                dp += matriz[i][j]
            }
        }
    }

    ds := 0
    for i := 0; i < 5; i++{
        ds += matriz[i][5-1-i]
    }

    result := dp - ds
    fmt.Println(result)
}

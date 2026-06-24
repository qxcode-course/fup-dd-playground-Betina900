package main
import "fmt"
func main() {
    var num, s int
    var matriz[2][3]int

    for i := 0; i < 2; i++{
        for j := 0; j < 3; j++{
            fmt.Scan(&num)
            matriz[i][j] = num
        }
    }

    for i := 0; i < 2; i++{
        for j := 0; j < 3; j++{
            s += matriz[i][j]
        }
    }

    fmt.Println(s)
}
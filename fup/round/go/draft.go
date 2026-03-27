package main
import (
    "fmt"
    "math"
)
func main() {
    var letra string
    var num float64
    fmt.Scan(&letra, &num)

    switch letra{
    case "c":
        fmt.Println(math.Ceil(num))
    case "f":
        fmt.Println(math.Floor(num))
    case "r":
        fmt.Println(math.Round(num))
    }
}

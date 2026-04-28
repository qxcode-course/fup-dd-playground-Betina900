package main
import "fmt"
func main() {
    var nums int
    fmt.Scan(&nums)
    arr := make([]int, nums)

    for i := 0; i<nums; i++{
        fmt.Scan(&arr[i])
    }

    fmt.Print("[")
    for i, elementos := range arr{
        if i == len(arr) - 1{
            fmt.Print(elementos)
        } else{
            fmt.Print(elementos, ", ")
        }
    }
    fmt.Println("]")
}

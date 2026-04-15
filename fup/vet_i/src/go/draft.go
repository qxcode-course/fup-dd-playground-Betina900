package main
import "fmt"
func main() {
    var elementos int
    fmt.Scan(&elementos)

    arr := make([]int, elementos)

    if elementos == 0 {
        fmt.Println()
        return
    }
    
    for i := 0; i < elementos; i++ {
        fmt.Scan(&arr[i])
    }

    for i := 0; i < elementos; i++ {
        fmt.Println(arr[i])
    }
}

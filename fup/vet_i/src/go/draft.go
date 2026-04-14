package main
import "fmt"
func main() {
    var elementos int
    fmt.Scan(&elementos)
    var arr []int = make([]int, elementos)
    for i:=0; i<len(arr); i++{
        fmt.Println(arr[i])
    }
}

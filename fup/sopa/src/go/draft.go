package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    if n == 0 || n == 1{
        fmt.Println(1)
        return
    }

    fib := make([]int64, n + 1)

    fib[0] = 1
    fib[1] = 1

    for i := 2; i <= n; i++{
        fib[i] = fib[i-1] + fib[i-2]
    }

    fmt.Println(fib[n - 1])
}
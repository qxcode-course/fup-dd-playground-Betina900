package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    for t := 0; t < n; t++{
        var numero string
        fmt.Scan(&numero)

        leds := 0
        led := []int{6, 2, 5, 5, 4, 5, 6, 3, 7,6}

        for i := 0; i < len(numero); i++{
            leds += led[numero[i] - '0']
        }

        fmt.Println(leds, "leds")
    }
}
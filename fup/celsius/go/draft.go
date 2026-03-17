package main
import "fmt"
func main() {
    var temp_celsius float64
    fmt.Scan(&temp_celsius)

    var temp_fahrenheit = (1.8*temp_celsius)+32
    fmt.Printf("%.6f\n", temp_fahrenheit)
}

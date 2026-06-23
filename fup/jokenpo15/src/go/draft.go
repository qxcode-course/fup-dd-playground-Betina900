package main
import "fmt"
func main() {
    var a, b int
    fmt.Scan(&a, &b)

    dist := (a - b + 15) % 15

    if dist == 0{
        fmt.Println("Empate")
    } else if dist <= 7{
        fmt.Println("Jogador 2")
    } else{
        fmt.Println("Jogador 1")
    }
}
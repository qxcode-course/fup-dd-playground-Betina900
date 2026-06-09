package main
import "fmt"
func main() {
    var jog1 int
    var jog2 int
    var jog3 int
    fmt.Scan(&jog1, &jog2, &jog3)

    if jog1 == 1 && jog2 == 1 && jog3 == 1|| jog1 == 0 && jog2 == 0 && jog3 == 0{
        fmt.Println("empate")
    } else if jog1 == 1 && jog2 == 0 && jog3 == 0 || jog1 == 0 && jog2 == 1 && jog3 == 1{
        fmt.Println("jog1")
    } else if jog1 == 0 && jog2 == 1 && jog3 == 0 || jog1 == 1 && jog2 == 0 && jog3 == 1{
        fmt.Println("jog2")
    } else if jog1 == 0 && jog2 == 0 && jog3 == 1 || jog1 == 1 && jog2 == 1 && jog3 == 0{
        fmt.Println("jog3")
    }
}
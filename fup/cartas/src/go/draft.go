package main
import "fmt"
func main() {
    var cartas_na_mao int
    fmt.Scan(&cartas_na_mao)

    fmt.Print("[")
    jogador := make([]string, cartas_na_mao)
    for i := 0; i < cartas_na_mao; i++ {
        fmt.Scan(&jogador[i])

        if jogador[i] == "1"{
            jogador[i] = "A"
        } else if jogador[i] == "11"{
            jogador[i] = "J"
        }else if jogador[i] == "12"{
            jogador[i] = "Q"
        } else if jogador[i] == "13"{
            jogador[i] = "K"
        }

        if i == cartas_na_mao - 1{
            fmt.Print(jogador[i])
        } else {
            fmt.Print(jogador[i], ", ")
        }
    } 
    fmt.Println("]")    
}

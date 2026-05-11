package main
import "fmt"
func abs(x int) int{
    if x<0{
        return -x
    }
    return x
}
func main() {
    var palpiteChico int
    var palpiteCebola int
    var quant_animais int
    fmt.Scan(&palpiteChico, &palpiteCebola, &quant_animais)

    letras := make([]string, quant_animais)
    for i:=0; i<quant_animais; i++{
        fmt.Scan(&letras[i])
    }

    patas := 0
   
    for _, animal := range letras{
        switch animal{
        case "v", "c":
            patas += 4
        case "g":
            patas += 2
        }
    }

    fmt.Println(patas)

    difChico := abs(patas - palpiteChico)
    difCebola := abs(patas - palpiteCebola)

    if difChico<difCebola{
        fmt.Println("Chico Bento")
    } else if difCebola<difChico{
        fmt.Println("Cebolinha")
    } else{
        fmt.Println("empate")
    }
}

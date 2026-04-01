package main
import "fmt"
func main() {
    var idade_mais_novo, quant_filhos int
    fmt.Scan(&idade_mais_novo, &quant_filhos)
    for i:=1; i<=quant_filhos; i++{
        fmt.Println(idade_mais_novo)
        idade_mais_novo+=2
    }
}

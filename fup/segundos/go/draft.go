package main
import "fmt"
func main() {
    var tempo int
    fmt.Scan(&tempo)
    var hora = tempo/3600
    var minutos = (tempo%3600)/60
    var segundo = tempo%60

    fmt.Printf("%d:%d:%d\n", hora, minutos, segundo)
}

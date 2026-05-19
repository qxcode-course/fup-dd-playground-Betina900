package main
import "fmt"
func inverter(numero int) int{
    invertido := 0

    for numero > 0{
        ultimoDigito := numero % 10

        invertido = invertido*10 + ultimoDigito

        numero = numero / 10
    }

    return invertido
}
func main() {
    var id int
    fmt.Scan(&id)

    if id == inverter(id){
        fmt.Println(1)
    } else{
        fmt.Println(0)
    }
}

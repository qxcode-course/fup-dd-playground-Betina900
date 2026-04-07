package main
import "fmt"
func sequenciaNumeroB(numeroA, numeroB int){
    for i:=numeroB; i>=numeroA; i--{
        fmt.Print(i, " ")
    }
}
func sequenciaNumeroA(numeroA, numeroB int){
    for i:=numeroA; i<=numeroB; i++{
        fmt.Print(i, " ", sequenciaNumeroB(numeroA, numeroB))
    }
}
func main() {
    var numeroA, numeroB int
    fmt.Scan(&numeroA, &numeroB)
    sequenciaCompleta(numeroA, numeroB)
}

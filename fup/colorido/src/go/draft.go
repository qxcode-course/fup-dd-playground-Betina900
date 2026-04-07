package main
import "fmt"
func main() {
    var n_caiu, contador int
    var peInicio, peDepois string
    fmt.Scan(&n_caiu, &peInicio)

    contador=0
    if peInicio=="d"{
        peDepois="e"
    } else{
        peDepois="d"
    }
    fmt.Print("[ ")
    for i:=0; i<=10; i++{
        if i==n_caiu{
            continue
        }
        if i==10{
            fmt.Print("ceu ")
            continue
        }
        if contador%2==0{
            fmt.Printf("%d%s ", i, peInicio)
        } else{
            fmt.Printf("%d%s ", i, peDepois)
        }
        contador++
    }
    fmt.Println("]")
}

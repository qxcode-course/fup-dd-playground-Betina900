package main
import "fmt"
func main() {
    var quantDindins, saborC, saborL, turnoM, turnoT int
    fmt.Scan(&quantDindins)
    
    sabor := make([]string, quantDindins)
    turno := make([]string, quantDindins)

    for i:=0; i<quantDindins; i++{
        fmt.Scan(&sabor[i])
        fmt.Scan(&turno[i])
    }

    for _, i := range sabor{
        if i=="c"{
            saborC+=1
        } else{
            saborL+=1
        }
    }
    for _, i := range turno{
        if i=="m"{
            turnoM+=1
        } else{
            turnoT+=1
        }
    }

    if saborC>saborL{
        fmt.Println("c")
    } else if saborC==saborL{
        fmt.Println("empate")
    } else{
        fmt.Println("l")
    }

    if turnoM==turnoT{
        fmt.Println("empate")
    } else if turnoM<turnoT{
        fmt.Println("m")
    } else{
        fmt.Println("t")
    }
}

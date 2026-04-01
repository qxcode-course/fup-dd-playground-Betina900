package main
import "fmt"
func main() {
    var dia_semana, hora int
    fmt.Scan(&dia_semana, &hora)

    if dia_semana>=2 && dia_semana<=6 {
        if hora>=8 && hora<=11{
            fmt.Println("SIM")
        } else if hora>=14 && hora<=17{
            fmt.Println("SIM")
        } else{
            fmt.Println("NAO")
        }
    }

    if dia_semana==7{
        if hora>=8 && hora<=11{
            fmt.Println("SIM")
        } else{
            fmt.Println("NAO")
        }
    }

    if dia_semana==1{
       fmt.Println("NAO")
    }
}

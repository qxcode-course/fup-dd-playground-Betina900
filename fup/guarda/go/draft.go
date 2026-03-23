 package main
import "fmt"
func main() {
    var wifi, login, adm bool
    fmt.Scan(&wifi, &login, &adm)

    1 := true
    0 := false
    switch {
    case wifi==0:
        fmt.Println("you must connect to wifi")
    case login==0:
        fmt.Println("you need to login first")
    case adm==0:
        fmt.Println("you must login as admin")
    case wifi==1 && login==1 && adm==1:
        fmt.Println("done")
    default:
        fmt.Println("ERRO")
    }
}

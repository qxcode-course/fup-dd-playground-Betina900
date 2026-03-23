 package main
import "fmt"
func main() {
    var wifi, login, adm int
    fmt.Scan(&wifi, &login, &adm)

    if wifi==0 {
        fmt.Println("you must connect to wifi")
    } else if login==0 {
        fmt.Println("you need to login first")
    } else if adm==0 {
        fmt.Println("you must to login as admin")
    } else {
        fmt.Println("done")
    }
}

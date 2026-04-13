package main
import "fmt"
func main() {
    var name string
    var year int
    fmt.Scan(&name, &year)

    if year<12{
        fmt.Println(name, "eh crianca")
    } else if year<18{
        fmt.Println(name, "eh jovem")
    } else if year<65{
        fmt.Println(name, "eh adulto")
    } else if year<1000{
        fmt.Println(name, "eh idoso")
    } else{
        fmt.Println(name, "eh mumia")
    }
}

package main
import "fmt"
func main() {
    var dim_tabu, horizonteX, verticalY, seg int
    var direcao string
    fmt.Scan(&dim_tabu, &horizonteX, &verticalY, &direcao, &seg)

    switch direcao{
    case "U":
        verticalY=(verticalY-seg+dim_tabu)%dim_tabu
        fmt.Println(horizonteX, verticalY)
    case "D":
        verticalY=(verticalY+seg+dim_tabu)%dim_tabu
        fmt.Println(horizonteX, verticalY)
    case "L":
        horizonteX=(horizonteX-seg+dim_tabu)%dim_tabu
        if horizonteX<0{
            fmt.Println(horizonteX*(-1), verticalY)
        } else{
            fmt.Println(horizonteX, verticalY)
        }
    case "R":
        horizonteX=(horizonteX+seg+dim_tabu)%dim_tabu
        fmt.Println(horizonteX, verticalY)
    }
}

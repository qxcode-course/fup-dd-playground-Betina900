package main
import "fmt"
func abs(x int) int{
    if x<0{
        return -x
    }
    return x
}

func main() {
    var valorReal int
    var chuteA int
    var chuteB int
    fmt.Scan(&valorReal, &chuteA, &chuteB)

    difA := abs(valorReal - chuteA)
    difB := abs(valorReal - chuteB)

    if difA < difB{
        fmt.Println("primeiro")
    } else if difB < difA{
        fmt.Println("segundo")
    } else{
        fmt.Println("empate")
    }    
}
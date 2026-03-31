package main
import "fmt"
func main() {
    var numA, numB int
    fmt.Scan(&numA, &numB)

    soma := 0

    if numB<numA{
        fmt.Println("invalido")
        return
    }
    
    for i:=numB; i>=numA; i--{
        if i%2==0{
            soma += i
        }
    }

    fmt.Println(soma)
}

package main
import "fmt"
func main() {
    var monica, filhoX, filhoY int
    fmt.Scan(&monica, &filhoX,&filhoY)

    filhoZ := monica - (filhoX+filhoY)
    
    if filhoX>filhoY && filhoX>filhoZ{
        fmt.Println(filhoX)
    } else if filhoY>filhoX && filhoY>filhoZ{
        fmt.Println(filhoY)
    } else{
        fmt.Println(filhoZ)
    }
}

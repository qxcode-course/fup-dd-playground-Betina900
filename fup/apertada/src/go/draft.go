package main
import "fmt"
func main() {
    var valores int
    fmt.Scan(&valores)
    nums := make([]int, valores) //criando lista

    for i:=0; i<valores; i++{ //lendo valores
        fmt.Scan(&nums[i])
    }
    menor := nums[0] //o menor número

    fmt.Println(menor)
}

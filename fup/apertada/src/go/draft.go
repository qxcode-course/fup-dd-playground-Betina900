package main
import "fmt"
func main() {
    nums := make([]int, 5) //criando lista de inteiros

    for i:=0; i<5; i++{ //lendo valores
        fmt.Scan(&nums[i])
    }
    menor := nums[0] //assume que o primeiro elemento é o menor
    
    for _, n := range nums{ //percorre o vetor
        if n<menor{
            menor = n
        }
    }
    fmt.Println(menor)
}

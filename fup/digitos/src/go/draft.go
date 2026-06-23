package main
import "fmt"
func main() {
    var a int //saber quantas vezes aparece (entre 1 e 9)
    var b int //número do contato
    fmt.Scan(&a)
    fmt.Scan(&b)

    if a == 0 && b == 0{
        fmt.Println(1)
        return
    }

    cont := 0

    for b > 0{
        digitos := b % 10

        if digitos == a{
            cont++
        }

        b /= 10
    }

    fmt.Println(cont)
}
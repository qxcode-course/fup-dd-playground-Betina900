package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    frase := scanner.Text()

    for i := len(frase)-1; i >= 0; i-- {
        fmt.Print(string(frase[i]))
    }

    fmt.Println()
}
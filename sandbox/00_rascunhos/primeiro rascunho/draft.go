package main

import "fmt"

func main() {
    var dia string = "seg"
    if dia == "seg" || dia == "ter" || dia == "qua"{
        fmt.Println("aula")
    } else if dia == "sab" || dia == "dom" {
        fmt.Println("final de semana")
    }

}
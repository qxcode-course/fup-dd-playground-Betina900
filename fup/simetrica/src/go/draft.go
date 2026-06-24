package main
import "fmt"
func main() {
    var transposta[3][3]int
    var matriz[3][3]int

    for i := 0; i < 3; i++{
        for j := 0; j < 3; j++{
            fmt.Scan(&matriz[i][j])
        }
    }

    for j := 0; j < 3; j++{
        for i := 0; i < 3; i++{
            transposta[j][i] = matriz[i][j]
        }
    }

    simetrik := true

    for i := 0; i < 3; i++{
        for j := 0; j < 3; j++{
            if matriz[i][j] != transposta[i][j]{
                simetrik = false
                break
            }
        }
    }

    if simetrik == true{
        fmt.Println("sim")
    } else{
        fmt.Println("nao")
    }
}
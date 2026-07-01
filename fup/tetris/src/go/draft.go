package main
import "fmt"
func main() {
    var L, C int
    fmt.Scan(&L, &C)

    tabuleiro := make([]string, L,)
    for i := 0; i < L; i++{
        fmt.Scan(&tabuleiro[i])
    }

    colidiu := false

    for i := 0; i < L; i++{
        for j := 0; j < C; j++{
            if tabuleiro[i][j] == 'o'{
                if i + 1 < L{
                    if tabuleiro[i+1][j] == '#'{
                        colidiu = true
                    }
                }
            }
        }
    }

    if colidiu{
        for i := 0; i < L; i++{
            fmt.Println(tabuleiro[i])
        }
        return
    }

    novo := make([][]byte, L)

    for  i := 0; i < L; i++{
        novo[i] = []byte(tabuleiro[i])
    }

    for i := L - 2; i >= 0; i--{
        for j := 0; j < C; j++{
            if novo[i][j] == 'o'{
                if novo[i + 1][j] == '.'{
                    novo[i][j] = '.'
                    novo[i+1][j] = 'o'
                }
            }
        }
    }

    for i := 0; i < L; i++{
        fmt.Println(string(novo[i]))
    }
}
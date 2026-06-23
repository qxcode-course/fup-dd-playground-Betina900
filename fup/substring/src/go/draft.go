package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var p string
	var c, q int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	p = scanner.Text()

	fmt.Scan(&c, &q)

	for q > 0{
        fmt.Print(string(p[c]))
    }
	fmt.Println()
}

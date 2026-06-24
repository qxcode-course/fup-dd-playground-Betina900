package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	texto := scanner.Text()

	var c, q int

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &c)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &q)

	if c < 0 || c >= len(texto) || q < 0 {
		fmt.Println("")
		return
	}

	fim := c + q

	if fim > len(texto) {
		fim = len(texto)
	}

	fmt.Println(texto[c:fim])
}
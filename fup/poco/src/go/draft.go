package main

import "fmt"

func main() {
	var profundidade int
	var forcaSalto int
	var escorrega int

	fmt.Scan(&profundidade)
	fmt.Scan(&forcaSalto)
	fmt.Scan(&escorrega)

	sapo := 0

	for sapo+forcaSalto < profundidade && sapo >= 0 {
		fmt.Println(sapo, sapo+forcaSalto)

		sapo = sapo + forcaSalto - escorrega

		if forcaSalto > 0 {
			forcaSalto -= 10
		}
	}

	if sapo+forcaSalto >= profundidade {
		fmt.Println(sapo, "saiu")
	} else {
		fmt.Println(sapo, "morreu")
	}
}
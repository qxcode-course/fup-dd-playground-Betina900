package main
import "fmt"
func main() {
	var pa, pg, ra, rg float64
	fmt.Scan(&pa, &pg, &ra, &rg)

	custoA := pa / ra
	custoB := pg / rg
	
	if custoA < custoB{
		fmt.Println("A")
	} else {
		fmt.Println("G")
	}
}

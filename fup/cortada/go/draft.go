package main
import "fmt"
func main() {
	var base, topo int
	fmt.Scan(&base, &topo)

	area_retangulo := 160*70
	metade_retangulo := area_retangulo/2
	area_trapezio := (base+topo)*70/2

	if area_trapezio>metade_retangulo{
		fmt.Println(1)
	} else if area_trapezio<metade_retangulo{
		fmt.Println(2)
	} else{
		fmt.Println(0)
	}
}

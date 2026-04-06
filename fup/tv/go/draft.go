package main
import "fmt"
func main() {
	var valor_tv float64
	var quant_parcelas int
	fmt.Scan(&valor_tv, &quant_parcelas)

	switch quant_parcelas{
	case 1:
		fmt.Printf("%.02f\n", valor_tv)
		fmt.Printf("%.02f\n", valor_tv)
	case 2:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.05))/2)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.05))
	case 3:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.1))/3)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.1))
	case 4:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.15))/4)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.15))
	case 5:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.2))/5)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.2))
	case 6:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.25))/6)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.25))
	case 7:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.3))/7)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.3))
	case 8:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.35))/8)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.35))
	case 9:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.4))/9)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.4))
	case 10:
		fmt.Printf("%.02f\n", (valor_tv+(valor_tv*0.45))/10)
		fmt.Printf("%.02f\n", valor_tv+(valor_tv*0.45))
	}
}

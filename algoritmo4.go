package main

import "fmt"

func main() {
	var numero1 float64
	fmt.Print("O seu primeiro número é ")
	fmt.Scan(&numero1)
	var numero2 float64
	fmt.Print("O seu segundo número é ")
	fmt.Scan(&numero2)
	var numero3 float64
	fmt.Print("O seu terceiro número é")
	fmt.Scan(&numero3)
	var resultado float64 = (numero1*2 + numero2*3 + numero3*5) / (3 + 2 + 5)
	fmt.Print("A média ponderada é ", resultado)
}

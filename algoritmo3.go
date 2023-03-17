package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("Qual é seu peso?")
	fmt.Scan(&peso)
	var altura float64
	fmt.Print("Qual é sua altura?")
	fmt.Scan(&altura)
	var resultado float64 = peso / (altura * altura)
	fmt.Print("O seu IMC é ", resultado)
}

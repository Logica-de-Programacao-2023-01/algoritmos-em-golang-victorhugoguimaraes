package main

import "fmt"

func main() {
	var numero1 uint
	fmt.Print("Qual é o número 1?")
	fmt.Scan(&numero1)
	var numero2 uint
	fmt.Print("Qual é o número 2?")
	fmt.Scan(&numero2)
	var numero3 uint
	fmt.Print("Qual é o numero 3?")
	fmt.Scan(&numero3)
	var resultado = (numero1 + numero2 + numero3)
	fmt.Print("O resultado é ", resultado)
}

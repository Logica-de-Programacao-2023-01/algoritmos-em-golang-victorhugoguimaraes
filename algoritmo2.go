package main

import "fmt"

func main() {
	var número uint
	fmt.Print("Qual é seu número?")
	fmt.Scan(&número)
	var dobro uint = (número * 2)
	var triplo uint = (número * 3)
	var quadruplo uint = (número * 4)
	fmt.Print("O dobro do teu número é ", dobro, ", o triplo do teu número é ", triplo, " e o quadruplo do seu número é ", quadruplo)
}

package main

import "fmt"

func main() {
	var idade uint
	fmt.Print("Qual é a sua idade? ")
	fmt.Scan(&idade)
	var resultado uint = idade * 365
	fmt.Print("A sua idade em dias é ", resultado)
}

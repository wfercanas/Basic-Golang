package main

import "fmt"

func main() {
	valor1 := 1
	valor2 := 2
	
	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}
	
	// With and
	if valor1 == 1 && valor2 == 3 {
		fmt.Println("Es verdad, &&")
	}

	// With or
	if valor1 == 1 || valor2 == 3 {
		fmt.Println("Es verdad, ||")
	}
}
package main

import "fmt"

func main() {
	modulo := 4 % 2
	switch modulo {
		case 0:
			fmt.Println("Es par")
		default:
			fmt.Println("Es impar")
	}
	
	// Sin condicion
	value := 200 
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es negativo")
	default:
		fmt.Print("Está entre cero y 100")
	}
}
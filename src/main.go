package main

import "fmt"

func main() {
	// Constantes
	const pi float64 = 3.1416
	const pi2 = 3.1416
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Variables
	base := 12
	var altura int = 14
	var area int

	area = base * altura
	fmt.Println("Area de un cuadrado:", area)

	// Zero Values
	var a int
	var b float32
	var c string
	var d bool
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
}
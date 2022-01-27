package main

import "fmt"

func constAndVar() {
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

func arithmetic() {
	x := 10
	y := 50
	sum := x + y
	less := x - y
	mult := x * y
	div := y / x
	mod := y % x
	fmt.Println("sum:", sum)
	fmt.Println("less:", less)
	fmt.Println("mult:", mult)
	fmt.Println("div:", div)
	fmt.Println("mod:", mod)
}

func fmtPackage() {
	// Println
	hello := "Hello"
	world := "World"
	fmt.Println(hello, world)
	fmt.Println(hello, world)
	
	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)
	
	// Sprintf
	message := fmt.Sprintf("Si!, %s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)
	
	// Data Type
	fmt.Printf("Hello: %T\n", hello)
	fmt.Printf("Cursos: %T", cursos)
}

func printFunction(text string) {
	fmt.Println(text)
}

func tripleParameters(a, b int, c string) {
	fmt.Println(a, b, c)
}

func withReturn(a int) int {
	return a * 2
}

func multipleReturn(a int) (c, d int) {
	return a*2, a*3
}

func functions() {
	printFunction("Hola mundo")
	printFunction("New string")
	tripleParameters(1, 2, "hola")
	fmt.Println("Single return:", withReturn(3))
	value1, value2 := multipleReturn(2)
	onlyValue1, _ := multipleReturn(3)
	fmt.Println("Multiple return:", value1, value2)
	fmt.Println("Only value:", onlyValue1)
}

func main() {
	// Class: const, var and zero-values
	// constAndVar()

	// Class: Arithmetic
	// arithmetic()

	// Class: fmt
	// fmtPackage()

	// Class: Functions
	functions()
}
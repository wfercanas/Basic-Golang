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

func forCycle() {
	// Conditional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	
	// While
	counter := 20
	for counter < 30 {
		fmt.Println(counter)
		counter++
	}
	
	// Forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 100000 {
			break
		}
	}
}

func conditionals() {
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

func switchConditional() {
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

func deferBreakContinue() {
	// defer
	fmt.Println("Hola")
	defer fmt.Println("Mundo")
	
	// continue and break
	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		if i == 8 {
			fmt.Println("Es 8")
			break
		}
		fmt.Println(i)
	}	
}

func arraysAndSlices() {
	// Array
	var array [4]int
	fmt.Println(array) // [0 0 0 0]
	
	array[0] = 1
	array[1] = 2
	fmt.Println(array) // [1 2 0 0]
	
	fmt.Println(array, len(array), cap(array)) // [1 2 0 0] 4 4
	
	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice)) // [0 1 2 3 4 5 6] 7 7
	
	// Métodos para slice
	fmt.Println(slice[0]) // 0
	fmt.Println(slice[:3]) // [0 1 2]
	fmt.Println(slice[2:4]) // [2 3]
	fmt.Println(slice[4:]) // [4 5 6]
	
	// Adicionar elementos (Append)
	slice = append(slice, 7)
	fmt.Println(slice) // [0 1 2 3 4 5 6 7]
	
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice) // [0 1 2 3 4 5 6 7 8 9 10]
}

func main() {
	// Class: const, var and zero-values
	// constAndVar()

	// Class: Arithmetic
	// arithmetic()

	// Class: fmt
	// fmtPackage()

	// Class: Functions
	// functions()

	// Class: For
	// forCycle()

	// Class: Conditionals
	// conditionals()
	// switchConditional()

	// Class: defer, break y continue
	// deferBreakContinue()

	// Class: Arrays and Slices
}
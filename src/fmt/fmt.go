package main

import "fmt"

func main() {
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
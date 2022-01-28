package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	square := cuadrado{base: 2}
	rectangle := rectangulo{base: 2, altura: 4}

	// Sin interfaz
	fmt.Println(square.area())    // 4
	fmt.Println(rectangle.area()) // 8

	// Con interfaz
	calcular(square)    // Area: 4
	calcular(rectangle) // Area: 8

	// Lista de interfaces
	myInterface := []interface{}{"Hola", 12, 4.902}
	fmt.Println(myInterface...)
}

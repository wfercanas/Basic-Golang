package main

import "fmt"

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

func main() {
	printFunction("Hola mundo")
	printFunction("New string")
	tripleParameters(1, 2, "hola")
	fmt.Println("Single return:", withReturn(3))
	value1, value2 := multipleReturn(2)
	onlyValue1, _ := multipleReturn(3)
	fmt.Println("Multiple return:", value1, value2)
	fmt.Println("Only value:", onlyValue1)
}
package main

import "fmt"

func main() {
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
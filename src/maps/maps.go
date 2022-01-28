package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["jose"] = 14
	m["pepito"] = 20

	fmt.Println(m)

	for i, value := range m {
		fmt.Println(i, value)
	}

	value := m["jose"]
	fmt.Println(value)

	wrong, ok := m["joseph"]
	fmt.Println(wrong, ok) // 0 false : porque no existe la llave
}
package main

import "fmt"

func main() {
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
package main

import "fmt"

func main() {
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
package main

import "fmt"

func isPalindromo(text string) bool {
	var textReversed string
	for i := len(text)-1; i >= 0; i-- {
		textReversed += string(text[i])
	}
	return text == textReversed
}

func main() {
	slice := []string{"Hola", "Que", "Hace"}
	
	// Both values	
	for i, valor := range slice {
		fmt.Println(i, valor)
	}
	
	// Only second return
	for _, valor := range slice {
		fmt.Println(valor)
	}
	
	// Only first return
	for i := range slice {
		fmt.Println(i)
	}

	text := "galleta"
	fmt.Println("Es palindromo la palabra " + text + "?", isPalindromo(text))
}
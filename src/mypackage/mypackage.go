package mypackage

import "fmt"

type carPrivate struct {
	brand string
	year 	int
}

// CarPublic with public access
type CarPublic struct {
	Brand string
	Year 	int
}

func PrintMessage() {
	fmt.Println("Hola")
}

func printPrivateMessage(text string) {
	fmt.Println(text)
}
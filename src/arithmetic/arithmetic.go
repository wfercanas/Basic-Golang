package main

import "fmt"

func main() {
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
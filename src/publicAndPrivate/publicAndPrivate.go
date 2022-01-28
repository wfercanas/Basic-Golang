package main

import (
	"fmt"
	"golang-platzi/src/mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	
	fmt.Println(myCar)
	
	// var myOtherCar mypackage.carPrivate
	// fmt.Println(myOtherCar)
	
	mypackage.PrintMessage()
}
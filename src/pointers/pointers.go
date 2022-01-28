package main

import "fmt"

type pc struct {
	ram int
	disk int
	brand string
}
func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong..")
}
func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(a) // 50
	fmt.Println(b) // 0xc00001c120
	fmt.Println(*b) // 50

	*b = 100
	fmt.Println(a) // 100

	myPc := pc{ram: 16, disk: 200, brand:"msi"}
	fmt.Println(myPc)
	myPc.ping()

	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)
}
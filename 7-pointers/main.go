package main

import "fmt"
//START OMIT
func one(xPtr *int) {
	*xPtr = 1
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
	fmt.Println(xPtr)//memory address of xPtr
}
//END OMIT
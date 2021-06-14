package main

import "fmt"

//START OMIT
type rectangle struct {
	length float64
	breadth float64

}
//Reciever functions
func (r* rectangle) setLength(newLength float64){
	r.length=newLength
}

func (r rectangle) areaCalculate() float64{
     return r.length*r.breadth
}
func main() {

	r:=rectangle{
		length:  34.5,
		breadth: 44.5,
	}
	fmt.Println("The area is ",r.areaCalculate())
    r.setLength(56.5)
	fmt.Println("The new area is ",r.areaCalculate())

}
//END OMIT
package main

import ("fmt"
"math")

type circle struct {
	x,y,r float64
}
func (c* circle) areaCircle() float64{

	return float64(math.Pi * c.r*c.r)
}


func main(){

	c:=circle{4,5,3}
	fmt.Println(c.x)
    fmt.Printf("The area is ",c.areaCircle())




}

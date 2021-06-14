package main

import (
	"fmt"
	"math"
	)
//START OMIT
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
 //returning multiple values
 func math_Cal (a float64)(float64,float64){

 	  s:=math.Sqrt(a)
 	  l:=math.Log(a)
 	  return s,l
}
//Variadic function
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

//END OMIT
//Closure
func makeOddGenerator () func ()int{
	i:=1
	return func() int {
		ret:=i
		i+=2
		return ret
	}

}
func main(){

	a:=[]float64{3.2,4.5,6.3,5.3}
    fmt.Println("The average is ", average(a))

	b:=49.0
	s,l:=math_Cal(b)
	fmt.Printf("The square root and log of %v are %v and %v\n", b,s,l)

	fmt.Println(add(1,2,3,6))

	gen:=makeOddGenerator()
	fmt.Println(gen())
	fmt.Println(gen())


}
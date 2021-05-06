package main

import ("fmt"

)
func main(){

	//initialise an array
	var a[5] int // will initialise all elements as 0
	fmt.Println(a)//[0 0 0 0 0]

	b:=[5]int{10,20,30,40,50} // initialise and declare together
	fmt.Println(b)//[10 20 30 40 50]

	// Declare and then initialise
    var c[5] int
	c[0]=10
	c[2]=30
	fmt.Println(c)//[10 0 30 0 0]

	// 2. Slice - dynamic size array
	// 2.1 Can be defined above
	var d []int
	fmt.Println("d empty:", d)//[]
	d = []int{10, 20, 30, 40}
	fmt.Println("d not empty:", d)// [10 20 30 40]


	// 2.2 Or like this - no need to specify size
	e := []int{10, 20, 30, 40, 50}
	fmt.Println("d:", e)//[10 20 30 40 50]


	// 2.3 Or like this
	s := make([]int, 4)
	s = append(s, 60, 70)
	fmt.Println("s:", s)//[0 0 0 0 60 70]

	// Manipulate slice

	s[2] = s[len(s)-1]
	fmt.Println("s:", s)
	fmt.Println("s(3-5)", s[2:5])

	// Range
	for k, v := range s {
		fmt.Printf("%d is %d\n", k, v)
	}
}
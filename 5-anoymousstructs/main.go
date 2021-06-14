package main

import "fmt"
//START OMIT
func main(){

	subjects:= struct {
		 name string
		 semester map[string]int
		 code int
	}{
		name:"Physical sensor",
		semester: map[string]int{
			"A":02,
			"B":03,

		},
		code: 123,

	}

	for key,value:= range subjects.semester{
		fmt.Printf("%v version :%v and semester %v\n",subjects.name,key,value)

	}
}
//END OMIT
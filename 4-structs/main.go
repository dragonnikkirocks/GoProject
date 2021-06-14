package main

import "fmt"

//START OMIT

type products struct {
	name string
    chara []string
}
func main(){

product1:= products{

	name:"Milk",
	chara: []string{
		"Bio ",
		"Protein",
	},
}

fmt.Println(product1.name)//dot operator to access each element
for _,value:= range product1.chara{
	fmt.Printf("%v is a charachteristic\n", value)
}
//END OMIT

type DiaryProducts struct {
	 products
	 vegan bool
}

DP1:=DiaryProducts{
	products:products{name:"Yougurt",
	chara: []string{
		"Bio","Easily available",
	},
	},
	vegan:false,
}
for _,value:= range DP1.chara{
	fmt.Printf("%v is a charachteristic\n", value)
}

}

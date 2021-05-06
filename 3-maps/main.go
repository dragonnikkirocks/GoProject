package main

import "fmt"

func main(){

	var m map[string]int//declaring
	m= make(map[string]int) //intialising
	fmt.Println(m)

	s:= map[string]string{
		"SEM1":"Chemistry",
		"SEM2":"Physics",
		"SEM3":"Mathematics",
	}

	//a. Adding to the map:
	m["SSTM34"]=01
	m["SSTM45"]=02
	fmt.Println(m)

	//b. Remove from the map:
	delete(m,"SSTM34")
	fmt.Println(m)

	m["SSTM06"]=05

	//c. ranging
	for key,value:= range s {
		fmt.Printf("key :%v and value :%v\n", key, value)
	}
}

package main

import (
	"fmt"
	"time"
)
func main(){

 var msg ="Hello"
 go func (){  //anoymous function

 	fmt.Println(msg) //msg will be accesible by closure
 					//nikki will get printed here as most of the time the go routine is not going to be
 					// until it hits the sleep call. Even though go routine is created, Nikki is printed
                 	//Race condition
 }()
 msg ="Nikki"
 time.Sleep(100 * time.Millisecond) // sleep calls are unreliable


}


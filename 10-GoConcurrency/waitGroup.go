package main

import (
	"fmt"
	"time"
	"sync"
)
func main(){

	var msg ="Hello"
	var wg = sync.Waitgroup{}
	wg.Add(1) // adding a go routine
	go func (msg string){  //anoymous function

		fmt.Println(msg)
		wg.Done() // telling that it's done
	}(msg)

	msg ="Nikki"
	wg.Wait()



}

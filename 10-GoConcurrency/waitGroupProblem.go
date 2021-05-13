package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter =0
func main(){

	for i:=0;i<10;i++ {
		wg.Add(2)
		go sayHello() //race condition
		go incrementCounter()

	}
	wg.Wait()

}

func sayHello(){

	fmt.Println("Hello there ",counter)
	wg.Done()
}
func incrementCounter(){

	counter++
	wg.Done()
}
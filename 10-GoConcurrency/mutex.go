package main

import( "fmt"
		"runtime"
		"sync"
      )

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var counter =0
func main(){

	runtime.GOMAXPROCS(100)
	for i:=0;i<10;i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go incrementCounter()

	}
	wg.Wait()

}

func sayHello(){

	fmt.Println("Hello there ",counter)
	m.RUnlock()
	wg.Done()
}
func incrementCounter(){
	counter++
	m.Unlock()
	wg.Done()
}



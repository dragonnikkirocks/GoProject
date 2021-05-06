package main

import("fmt")

func main(){

//single for loop
sum:=0
var i=0
for i:=1;i<5;i++{
  sum+=i
}
  fmt.Println(sum)

//single while loop
for i<12{

  i++
}
fmt.Println(i)
//infinte loop
for{
  sum++

}
}

package main

import "fmt"

func main(){
	var a = 1
	var b = "initial"
	fmt.Println(b,a)

	var c,d int = 4,5
	fmt.Println(c,d)

	f:= "short"
	fmt.Println(f) //The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "short"

	var e int //by default will get value 0
	fmt.Println(e)

	var t=true //dynamic typing
	fmt.Println(t) 
}
package main

import "fmt"

//DEFINE GLOBAL VARIABLE
var a = 10

func main() {
	//ACCESS GLOBAL VARIABLE
	fmt.Println("FROM main() VALUE OF GLOBAL VARIABLE A : ", a)
	a++
	//ACCESS GLOBAL VARIABLE FROM OTHER FUNCTION
	display()
}

func display() {
	//ACCESS GLOBAL VARIABLE
	fmt.Println("FROM display() VALUE OF GLOBAL VARIABLE A  : ", a)
}

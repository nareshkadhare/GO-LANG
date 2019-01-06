package main

import "fmt"

func main() {

	var a = 10
	var b = 20

	//function calling
	var sum = sum(a, b)
	fmt.Printf("SUM OF %d +  %d = %d ", a, b, sum)
}

//function defination
//func function_name( [parameter list] ) [return_types]
func sum(a int, b int) int {
	return a + b
}

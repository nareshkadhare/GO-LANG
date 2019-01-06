package main

import "fmt"

func main() {

	var a = 10
	var b = 20

	//function calling
	var sum = sum(a, b)
	fmt.Printf("SUM OF %d +  %d = %d ", a, b, sum)

	fmt.Printf("\n Before Swaping  A : %d , B : %d ", a, b)
	
	//function call handle multiple values
	a, b = swap(a, b)
	fmt.Printf("\n After Swaping  A : %d , B : %d ", a, b)
}

//function defination
//func function_name( [parameter list] ) [return_types]
func sum(a int, b int) int {
	return a + b
}

//function can return multiple values
func swap(a int, b int) (int, int) {
	return b, a
}

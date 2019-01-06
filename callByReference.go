package main

import "fmt"

func main() {

	var a = 10
	var b = 20

	fmt.Printf("\n Before Swaping  A : %d , B : %d ", a, b)

	// call by reference Used & symbol To Pass Reference
	swap(&a, &b)

	fmt.Printf("\n After Swaping  A : %d , B : %d ", a, b)
}

// function used Pointer type to handle Reference Values
func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

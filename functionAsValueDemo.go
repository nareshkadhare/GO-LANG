package main

import "fmt"

func main() {
	getSqure := func(x int) int {
		return (x * x)
	}

	fmt.Println(" SQUERE : ", getSqure(8))
}

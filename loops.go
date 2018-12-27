package main

import "fmt"

func main() {
	var n = 5
	fmt.Println("-------SIMPLE FOR LOOP LIKE C LANG----------")
	//SIMPLE FOR LOOP LIKE C LANG
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n-------FOR LOOP USED LIKE WHILE LOOP----------")
	//FOR LOOP USED LIKE WHILE LOOP
	for n <= 10 {
		fmt.Println(n)
		n++
	}

	fmt.Println("\n-------FOR EACH LOOP----------")
	var listVal = [5]int{10, 20, 30, 40, 50}
	for i, val := range listVal {
		fmt.Println("INDEX : ", i, " == >", val)
	}

}

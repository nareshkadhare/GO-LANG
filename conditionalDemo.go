package main

import "fmt"

func main() {

	var a = 30
	var b = 20

	//simple if structure
	if a > b {
		fmt.Println(a, "IS GRETER THAN", b)
	}

	a = 10
	b = 15
	//simple if else structure
	if a > b {
		fmt.Println(a, "IS GRETER THAN", b)
	} else {
		fmt.Println(b, "IS GRETER THAN", a)
	}

	a = 10
	b = 15
	//if else if structure
	if a > b {
		fmt.Println(a, "IS GRETER THAN", b)
	} else if b > a {
		fmt.Println(b, "IS GRETER THAN", a)
	} else {
		fmt.Println(b, " ARE EQUAL ", a)
	}

	a = 60
	b = 120
	c := 190
	//nested if
	if a > b {
		if a > c {
			fmt.Println(a, "IS GRETER THAN", b, ",", c)
		} else {
			fmt.Println(c, "IS GRETER THAN", a)
		}
	} else if b > c {
		fmt.Println(b, "IS GRETER THAN", a, ",", c)
	} else {
		fmt.Println(c, "IS GRETER THAN", a, ",", b)
	}
}

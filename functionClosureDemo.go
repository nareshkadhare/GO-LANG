package main

import "fmt"

func main() {

	//function as Variable Which Return Function
	autoIncrement := func() func() int {
		var i = 0
		//anonymous function define
		return func() int {
			i++
			return i
		}
	}

	//Get anonymous function
	numSequence := autoIncrement()

	//Call numSequence wich generate int values on each call eg. first call : 1 second Call : 2
	fmt.Println(" 1st CALL : ", numSequence())
	fmt.Println(" 2nd CALL : ", numSequence())
	fmt.Println(" 3rd CALL : ", numSequence())
	fmt.Println(" 4th CALL : ", numSequence())
	fmt.Println(" 5th CALL : ", numSequence())
}

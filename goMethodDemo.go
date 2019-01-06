package main

import "fmt"

// Define Student Structure
type Student struct {
	rno  int
	name string
}

//Define Methode Wich Associeat To Student Type Variables
func (student Student) display() {
	fmt.Println("ROLL NO : ", student.rno)
	fmt.Println("NAME : ", student.name)
}

func main() {

	//Create Student Type
	dharmesh := Student{rno: 1, name: "Dharmesh"}
	//Call Define Method
	dharmesh.display()

	//Create Student Type
	amar := Student{rno: 2, name: "Amar"}
	//Call Define Method
	amar.display()
}

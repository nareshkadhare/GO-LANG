package main

import "fmt"

func main() {

	var a = 20
	//Integral Type
	// NOTE : SWITCH VARIABLE DATA TYPE MUST MATCH WITH CASE VALUE DATATYPE OR ITS LOWER DATA TYPE
	switch a {
	case 10:
		fmt.Println("VALUE IS 10")
		break
	case 20:
		fmt.Println("VALUE IS 20")
		break
	default:
		fmt.Println("DEFAUL BLOCK")
	}

	//Expression Switch
	a = 30
	switch a {
	case 5 + 5:
		fmt.Println("VALUE IS SUM OF 5 + 5 = ", a)
		break
	case 10 + 20:
		fmt.Println("VALUE IS SUM OF 10 + 20 = ", a)
		break
	case 10 * 5:
		fmt.Println("VALUE IS SUM OF 10 * 5 = ", a)
		break
	default:
		fmt.Println("DEFAUL BLOCK")
	}

	//with String Data Type
	grade := "A+"
	switch grade {
	case "A+":
		fmt.Println("Excellent")
		break
	case "A":
		fmt.Println("DIST")
		break

	case "B":
		fmt.Println("FIRST CLASS")
		break

	case "C":
		fmt.Println("SECOND CLASS")
		break

	default:
		fmt.Println("PASS CLASS")
	}

	//TYPE SWITCH
	//NOTE : The expression used in a switch statement must have an variable of interface{} type.
	var d interface{}
	d = 20.36
	switch v := d.(type) {
	case nil:
		fmt.Println("NIL", v)
	case float64:
		fmt.Println("FLOAT 64 : ", v)
	case int64:
		fmt.Println("INT 64 : ", v)
	default:
		fmt.Println("TYPE : ", v)
	}
}

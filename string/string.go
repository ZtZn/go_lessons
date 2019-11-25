package main

import "strconv"

import "fmt"

const (
	PI = 3.14
	SI = "Si"
	A  = iota
	B  = iota * 10
	C
	D
	E
)

func main() {
	// str := "String"
	// s := `String`
	// a:= 'r'  //rouna OR char

	//string to int

	num, error := strconv.Atoi("10a")
	if error != nil {
		fmt.Println(error.Error())
	}

	fmt.Println(num)

	fmt.Println(strconv.Itoa(10))

	number,_ := strconv.Atoi("100")
	fmt.Println(number)
}

package main

import "fmt"

func main() {
	a := []int64{1, 2, 3, 4, 5}
	for _, as := range a {
		fmt.Println(as)
	}
	var b = make([]int, 5)

	for range b {
		fmt.Println("Hello World")
	}

	//last element
	fmt.Println(a[len(a)-1])

	//[1 2 3]
	fmt.Println(a[0:3])
}

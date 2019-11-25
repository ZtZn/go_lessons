package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("my favorite number is", rand.Intn(10))
	fmt.Println(add(2, 3))
}

func add(x, y int) int {
	return x + y

}

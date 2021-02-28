package main

import "fmt"

type meutipo int

var x meutipo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42

	fmt.Println(x)
}

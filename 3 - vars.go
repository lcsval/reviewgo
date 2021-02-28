package main

import "fmt"

var g = 3

func main() {
	x := 10
	y := "olá bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)

	h := 65
	qualquercoisa(h)
}

func qualquercoisa(x int) {
	fmt.Println(g)
	fmt.Println(x)
}

package main

import "fmt"

func main() {
	x := 15
	a := &x // memory address of x
	fmt.Println(a)
	fmt.Println(*a)

	*a = 2
	fmt.Println(x) // value of x was changed using address
}
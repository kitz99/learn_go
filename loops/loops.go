package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	x := 5
	for {
		fmt.Println("do stuff", x)
		x += 3
		if x > 25 {
			break
		}
	}

}

package main

import (
	"fmt"
)

func main() {
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

	fmt.Println(grades)
	fmt.Println()
	delete(grades, "Timmy")
	fmt.Println(grades)
	fmt.Println()

	for k, v := range grades {
		fmt.Printf("%s -> %f\n", k, v)
	}

}

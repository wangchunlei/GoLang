package main

import (
		"fmt"
		"wangcl/newmath"
)

func main() {
	fmt.Printf("Hello, world.  Sqrt(2) = %v\n", newmath.Sqrt(2))
	for i := 0; i < 5; i++ {
    	defer fmt.Printf("%d ", i)
	}
}
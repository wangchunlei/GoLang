package main

import (
	"fmt"
)

type bigStruct struct {
	lots [1e6]float64
}

func main() {
	t := make([]map[string]bigStruct, 1e6)
	i := 0
	for {
		t[i] = make(map[string]bigStruct, 1000)
		fmt.Println("aaaaaaaaaaaaaaaaaaa")
		i = i + 1
	}
}

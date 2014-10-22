package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for i, addr := range addrs {
		fmt.Printf("%d %v\n", i, addr)
	}
	a := "wangchunlei"
	fmt.Println(*&a)
}

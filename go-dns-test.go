package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	hostname := os.Args[1]
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		fmt.Printf("found: %s\n", addr)
	}
}

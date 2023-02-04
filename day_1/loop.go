package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n = rand.Intn(10)

	if n > 4 {
		fmt.Printf("success")
	} else {
		println("failed")
	}
}

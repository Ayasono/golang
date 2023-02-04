package main

import "fmt"

func main() {
	const distance = 56000000
	const arrivalDays = 28

	var speed = distance / arrivalDays

	fmt.Printf("%v", speed)
}

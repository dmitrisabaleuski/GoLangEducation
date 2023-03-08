package main

import "fmt"

func main() {
	type celsius int
	const degrees = 20
	var temperature celsius = degrees

	temperature += 10

	fmt.Println(temperature)
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Мой вес на поверхности Марса равен ")
	fmt.Print(55.0 * 0.3783)
	fmt.Print(" килограммам, а мой возраст равен ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" годам.")

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

	const lightSpeed = 299792
	var distance = 56000000
	//var distance, speed = 56000000, 100800

	fmt.Println(distance/lightSpeed, "секунд")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "секунд")

	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}

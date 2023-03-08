package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64
type kelvin float64

func main() {
	const degrees = 20
	var temperature celsius = degrees
	var far fahrenheit = 20
	var warmUp float64 = 10
	temperature += celsius(warmUp)
	fmt.Println(temperature, far)

	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(k, "° K is ", c, "° C")

	var cel celsius = 127.0
	kel := celsiusToKelvin(cel)
	fmt.Println(cel, "° C is ", kel, "° K")

	fmt.Println(k.celsius())
	fmt.Println(far.celsius())
	fmt.Println(cel.fahrenheit())
	fmt.Println(k.fahrenheit())
	fmt.Println(cel.kelvin())
	fmt.Println(far.kelvin())
}

func kelvinToCelsius(k kelvin) celsius { //Function
	return celsius(k - 273.15)
}

func celsiusToKelvin(c celsius) kelvin { //Function
	return kelvin(c + 273.15)
}

func (k kelvin) celsius() celsius { //Method
	return celsius(k - 273.15)
}

func (f fahrenheit) celsius() celsius { //Method
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func (k kelvin) fahrenheit() fahrenheit { //Method
	return k.celsius().fahrenheit()
}

func (c celsius) fahrenheit() fahrenheit { //Method
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func (c celsius) kelvin() kelvin { //Method
	return kelvin(c + 273.15)
}

func (f fahrenheit) kelvin() kelvin { //Method
	return f.celsius().kelvin()
}

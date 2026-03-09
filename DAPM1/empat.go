package main

import "fmt"

func main() {
	var c, f, r, k float64

	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&c)

	r = c * (4.0 / 5.0)
	f = (c * (9.0 / 5.0)) + 32
	k = (f + 459.67) * (5.0 / 9.0)

	fmt.Printf("Derajat Reamur: %.0f\n", r)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", f)
	fmt.Printf("Derajat Kelvin: %.0f\n", k)
}
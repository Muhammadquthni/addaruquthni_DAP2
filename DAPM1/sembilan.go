package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	var akar2 float64 = 1.0

	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	for i := 0; i <= k; i++ {
		pembilang := math.Pow(float64(4*i+2), 2)
		penyebut := float64((4*i + 1) * (4*i + 3))
		akar2 = akar2 * (pembilang / penyebut)
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", akar2)
}
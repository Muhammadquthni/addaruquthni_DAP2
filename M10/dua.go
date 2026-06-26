package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var ikan [1000]float64
	totalSemua := 0.0

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
		totalSemua += ikan[i]
	}

	var wadah []float64
	var beratWadah float64
	count := 0

	for i := 0; i < x; i++ {
		beratWadah += ikan[i]
		count++

		if count == y || i == x-1 {
			wadah = append(wadah, beratWadah)
			beratWadah = 0
			count = 0
		}
	}

	for i := 0; i < len(wadah); i++ {
		fmt.Printf("%.2f", wadah[i])
		if i < len(wadah)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	if len(wadah) > 0 {
		rataRata := totalSemua / float64(len(wadah))
		fmt.Printf("%.2f\n", rataRata)
	}
}
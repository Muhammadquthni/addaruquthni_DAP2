package main

import "fmt"

func main() {
	var berat, kg, sisa, biayaKg, biayaSisa, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	kg = berat / 1000
	sisa = berat % 1000

	biayaKg = kg * 10000

	if kg >= 10 {
		biayaSisa = 0
	} else {
		if sisa >= 500 {
			biayaSisa = sisa * 5
		} else {
			biayaSisa = sisa * 15
		}
	}

	total = biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp. %d\n", total)
}
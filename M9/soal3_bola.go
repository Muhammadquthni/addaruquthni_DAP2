package main

import (
	"fmt"
)

const NMAX = 100

func main() {
	var klubA, klubB string
	var hasil [NMAX]string
	var jumlah int

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for jumlah < NMAX {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil[jumlah] = klubA
		} else if skorB > skorA {
			hasil[jumlah] = klubB
		} else {
			hasil[jumlah] = "Draw"
		}
		jumlah++
		pertandingan++
	}

	for i := 0; i < jumlah; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}

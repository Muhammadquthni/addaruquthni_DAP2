package main

import "fmt"

func main() {
	var t1, t2, t3, t4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scanln(&t1, &t2, &t3, &t4)

		if t1 != "merah" || t2 != "kuning" || t3 != "hijau" || t4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Printf("BERHASIL: %t\n", berhasil)
}
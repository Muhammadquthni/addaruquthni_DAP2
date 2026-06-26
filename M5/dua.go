package main

import "fmt"

func cetakSatuBaris(banyak int) {
	if banyak == 0 {
		return
	}
	fmt.Print("*")
	cetakSatuBaris(banyak - 1)
}

func polaBintang(n, sekarang int) {
	if sekarang > n {
		return
	}
	cetakSatuBaris(sekarang)
	fmt.Println()
	polaBintang(n, sekarang+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	polaBintang(n, 1)
}
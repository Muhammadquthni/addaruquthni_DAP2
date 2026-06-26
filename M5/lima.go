package main

import "fmt"

func cetakGanjil(n, sekarang int) {
	if sekarang > n {
		fmt.Println()
		return
	}
	if sekarang%2 != 0 {
		fmt.Print(sekarang, " ")
	}
	cetakGanjil(n, sekarang+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakGanjil(n, 1)
}
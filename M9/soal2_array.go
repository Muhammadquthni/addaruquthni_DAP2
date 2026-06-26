package main

import (
	"fmt"
	"math"
)

const NMAX = 100

func tampilSemua(arr [NMAX]int, n int) {
	fmt.Print("Semua elemen: ")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilGanjil(arr [NMAX]int, n int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilGenap(arr [NMAX]int, n int) {
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func tampilKelipatanX(arr [NMAX]int, n int, x int) {
	fmt.Print("Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func hapusElemen(arr *[NMAX]int, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}
	*n--
}

func rataRata(arr [NMAX]int, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return float64(total) / float64(n)
}

func standarDeviasi(arr [NMAX]int, n int) float64 {
	rata := rataRata(arr, n)
	total := 0.0
	for i := 0; i < n; i++ {
		total += (float64(arr[i]) - rata) * (float64(arr[i]) - rata)
	}
	return math.Sqrt(total / float64(n))
}

func frekuensi(arr [NMAX]int, n int, bil int) int {
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == bil {
			count++
		}
	}
	return count
}

func main() {
	var arr [NMAX]int
	var n int

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan", n, "bilangan:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	tampilSemua(arr, n)
	tampilGanjil(arr, n)
	tampilGenap(arr, n)

	var x int
	fmt.Print("Masukkan nilai kelipatan x: ")
	fmt.Scan(&x)
	tampilKelipatanX(arr, n, x)

	var idx int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)
	hapusElemen(&arr, &n, idx)
	fmt.Print("Array setelah hapus: ")
	tampilSemua(arr, n)

	fmt.Printf("Rata-rata: %.2f\n", rataRata(arr, n))
	fmt.Printf("Standar deviasi: %.2f\n", standarDeviasi(arr, n))

	var bil int
	fmt.Print("Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&bil)
	fmt.Println("Frekuensi", bil, ":", frekuensi(arr, n, bil))
}

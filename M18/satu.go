package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pita string
var indeks int
var karakterSekarang byte

func start(input string) {
	pita = input
	indeks = 0
	if len(pita) > 0 {
		karakterSekarang = pita[indeks]
	}
}

func maju() {
	indeks++
	if indeks < len(pita) {
		karakterSekarang = pita[indeks]
	}
}

func eop() bool {
	return karakterSekarang == '.'
}

func cc() byte {
	return karakterSekarang
}

func main() {
	fmt.Print("Masukkan teks (akhiri dengan tanda titik '.'): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	teks := scanner.Text()
	teks = strings.ToUpper(teks)

	start(teks)

	jumlahKarakter := 0
	jumlahA := 0
	jumlahLE := 0
	karakterSebelumnya := byte(' ')

	for !eop() && indeks < len(pita) {
		k := cc()

		jumlahKarakter++

		if k == 'A' {
			jumlahA++
		}

		if karakterSebelumnya == 'L' && k == 'E' {
			jumlahLE++
		}

		karakterSebelumnya = k
		maju()
	}

	fmt.Printf("\n--- HASIL PEMBACAAN MESIN ---\n")
	fmt.Printf("Pita selesai dibaca.\n")
	fmt.Printf("Jumlah karakter (sebelum titik): %d\n", jumlahKarakter)
	fmt.Printf("Jumlah huruf A: %d\n", jumlahA)

	frekuensiA := 0.0
	if jumlahKarakter > 0 {
		frekuensiA = float64(jumlahA) / float64(jumlahKarakter)
	}
	fmt.Printf("Frekuensi huruf A: %.2f\n", frekuensiA)

	fmt.Printf("Banyak kata LE: %d\n", jumlahLE)
}
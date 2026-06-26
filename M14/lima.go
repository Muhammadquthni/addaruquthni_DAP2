package main

import "fmt"

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku struct {
	pustaka  [7919]Buku
	nPustaka int
}

func DaftarkanBuku(p *DaftarBuku, n int) {
	p.nPustaka = n
	for i := 0; i < n; i++ {
		fmt.Scan(&p.pustaka[i].id, &p.pustaka[i].judul, &p.pustaka[i].penulis, &p.pustaka[i].penerbit, &p.pustaka[i].eksemplar, &p.pustaka[i].tahun, &p.pustaka[i].rating)
	}
}

func CetakTerfavorit(p DaftarBuku, n int) {
	if n == 0 {
		return
	}
	maxIdx := 0
	for i := 1; i < n; i++ {
		if p.pustaka[i].rating > p.pustaka[maxIdx].rating {
			maxIdx = i
		}
	}
	fmt.Println(p.pustaka[maxIdx].judul, p.pustaka[maxIdx].penulis, p.pustaka[maxIdx].penerbit, p.pustaka[maxIdx].tahun)
}

func UrutBuku(p *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := p.pustaka[i]
		j := i
		for j > 0 && temp.rating > p.pustaka[j-1].rating {
			p.pustaka[j] = p.pustaka[j-1]
			j--
		}
		p.pustaka[j] = temp
	}
}

func Cetak5Terbaru(p DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Println(p.pustaka[i].judul)
	}
}

func CariBuku(p DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1
	ketemu := false

	for kiri <= kanan && !ketemu {
		tengah := (kiri + kanan) / 2
		if p.pustaka[tengah].rating == r {
			fmt.Println(p.pustaka[tengah].judul, p.pustaka[tengah].penulis, p.pustaka[tengah].penerbit, p.pustaka[tengah].tahun, p.pustaka[tengah].eksemplar, p.pustaka[tengah].rating)
			ketemu = true
		} else if p.pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var db DaftarBuku
	var n, ratingCari int

	fmt.Scan(&n)
	DaftarkanBuku(&db, n)
	
	CetakTerfavorit(db, n)
	UrutBuku(&db, n)
	Cetak5Terbaru(db, n)
	
	fmt.Scan(&ratingCari)
	CariBuku(db, n, ratingCari)
}
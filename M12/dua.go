package main

import "fmt"

func main() {
	var suara [21]int
	var totalSuara, suaraSah, val int

	for {
		fmt.Scan(&val)
		if val == 0 {
			break
		}
		
		totalSuara++
		
		if val >= 1 && val <= 20 {
			suaraSah++
			suara[val]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuara)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	var ketua, wakil int
	var maxKetua, maxWakil int

	for i := 1; i <= 20; i++ {
		if suara[i] > maxKetua {
			maxWakil = maxKetua
			wakil = ketua
			
			maxKetua = suara[i]
			ketua = i
		} else if suara[i] > maxWakil {
			maxWakil = suara[i]
			wakil = i
		}
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}
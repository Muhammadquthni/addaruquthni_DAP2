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

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
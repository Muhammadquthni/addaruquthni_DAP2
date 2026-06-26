package main

import "fmt"

func main() {
	var dat, max int
	const MARKER = -9999

	fmt.Scan(&dat)

	if dat == MARKER {
		fmt.Println("tidak ada data")
	} else {
		max = dat
		fmt.Scan(&dat)

		for dat != MARKER {
			if dat > max {
				max = dat
			}
			fmt.Scan(&dat)
		}

		fmt.Println("Data terbesar", max)
	}
}
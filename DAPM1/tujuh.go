package main

import "fmt"

func main() {
	var bunga, pita string
	var count int
	
	count = 0
	for {
		fmt.Printf("Bunga %d: ", count+1)
		fmt.Scanln(&bunga)

		if bunga == "SELESAI" {
			break
		}

		if count == 0 {
			pita = bunga
		} else {
			pita = pita + " - " + bunga
		}
		count++
	}

	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %d\n", count)
}

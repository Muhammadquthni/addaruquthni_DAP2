package main

import "fmt"

func main() {
	var b int
	var totalFaktor int = 0

	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
			totalFaktor++
		}
	}
	fmt.Println()

	var prima bool = (totalFaktor == 2)
	fmt.Printf("Prima: %t\n", prima)
}
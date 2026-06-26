package main

import "fmt"

func sortAscending(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		idxMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func sortDescending(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		idxMax := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var ganjil, genap []int
		for j := 0; j < m; j++ {
			var val int
			fmt.Scan(&val)
			if val%2 != 0 {
				ganjil = append(ganjil, val)
			} else {
				genap = append(genap, val)
			}
		}

		sortAscending(ganjil)
		sortDescending(genap)

		for j := 0; j < len(ganjil); j++ {
			fmt.Print(ganjil[j], " ")
		}
		for j := 0; j < len(genap); j++ {
			fmt.Print(genap[j], " ")
		}
		fmt.Println()
	}
}
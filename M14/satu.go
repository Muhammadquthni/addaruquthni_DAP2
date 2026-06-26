package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		for j := 0; j < m-1; j++ {
			idxMin := j
			for k := j + 1; k < m; k++ {
				if arr[k] < arr[idxMin] {
					idxMin = k
				}
			}
			temp := arr[j]
			arr[j] = arr[idxMin]
			arr[idxMin] = temp
		}

		for j := 0; j < m; j++ {
			fmt.Print(arr[j])
			if j < m-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
package main

import "fmt"

func main() {
	var arr []int
	var val int

	for {
		fmt.Scan(&val)
		if val < 0 {
			break
		}
		arr = append(arr, val)
	}

	n := len(arr)
	for i := 1; i < n; i++ {
		temp := arr[i]
		j := i
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	if n < 2 {
		fmt.Println("Data berjarak tidak tetap")
	} else {
		jarak := arr[1] - arr[0]
		tetap := true
		for i := 2; i < n; i++ {
			if arr[i]-arr[i-1] != jarak {
				tetap = false
				break
			}
		}

		if tetap {
			fmt.Printf("Data berjarak %d\n", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}
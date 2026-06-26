package main

import "fmt"

func main() {
	var data []int
	var val int

	for {
		fmt.Scan(&val)

		if val == -5313 {
			break
		}

		if val == 0 {
			n := len(data)
			for i := 1; i < n; i++ {
				temp := data[i]
				j := i
				for j > 0 && temp < data[j-1] {
					data[j] = data[j-1]
					j--
				}
				data[j] = temp
			}

			if n > 0 {
				if n%2 != 0 {
					fmt.Println(data[n/2])
				} else {
					median := (data[(n/2)-1] + data[n/2]) / 2
					fmt.Println(median)
				}
			}
		} else {
			data = append(data, val)
		}
	}
}
package main

import "fmt"

func cariFaktor(n, pembagi int) {
	if pembagi > n {
		fmt.Println()
		return
	}
	if n%pembagi == 0 {
		fmt.Print(pembagi, " ")
	}
	cariFaktor(n, pembagi+1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cariFaktor(n, 1)
}
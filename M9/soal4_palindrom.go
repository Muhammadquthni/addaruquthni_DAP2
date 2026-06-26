package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	*n = 0
	for _, ch := range line {
		if ch == '.' {
			break
		}
		if ch != ' ' && *n < NMAX {
			t[*n] = ch
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var salinan tabel
	for i := 0; i < n; i++ {
		salinan[i] = t[i]
	}
	balikanArray(&salinan, n)
	for i := 0; i < n; i++ {
		if t[i] != salinan[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks : ")
	isiArray(&tab, &m)

	isPalin := palindrom(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)

	fmt.Printf("Palindrom ? %v\n", isPalin)
}

package main

import "fmt"

func main() {
	nilaiUjian := 96

	fmt.Printf("Nilai ujian Anda adalah: %d\n", nilaiUjian)

	if nilaiUjian >= 90 {
		fmt.Println("Predikat: A (Luar Biasa!)")
	} else if nilaiUjian >= 80 {
		fmt.Println("Predikat: B (Bagus)")
	} else if nilaiUjian >= 70 {
		fmt.Println("Predikat: C (Cukup)")
	} else {
		fmt.Println("Predikat: D (Perlu belajar lagi)")
	}

	fmt.Println("==================================")

	angka := -1

	fmt.Printf("Angka adalah: %d\n", angka)

	if angka > 0 {
		fmt.Printf("%d adalah angka positif.\n", angka)
	} else if angka < 0 {
		fmt.Printf("%d adalah angka negatif.\n", angka)
	} else {
		fmt.Printf("Angka adalah nol.\n")
	}
}

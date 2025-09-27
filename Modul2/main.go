package main

import "fmt"

func main() {
	char := 'z'

	switch char {
	case 'a', 'i', 'u', 'e', 'o':
		fmt.Printf("'%c' adalah huruf vokal.\n", char)
	default:
		fmt.Printf("'%c' adalah huruf konsonan.\n", char)
	}

	fmt.Println("==================================")

	nilaiUjian := 68

	fmt.Printf("Nilai ujian Anda: %d\n", nilaiUjian)

	switch {
	case nilaiUjian >= 90:
		fmt.Println("Predikat: A (Luar Biasa!)")
	case nilaiUjian >= 80:
		fmt.Println("Predikat: B (Bagus)")
	case nilaiUjian >= 70:
		fmt.Println("Predikat: C (Cukup)")
	default:
		fmt.Println("Predikat: D (Perlu belajar lagi)")
	}

	fmt.Println("==================================")

	nomorHari := 8

	fmt.Printf("Nomor hari: %d\n", nomorHari)

	switch nomorHari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Nomor hari tidak valid")
	}
}

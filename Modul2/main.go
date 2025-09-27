package main

import "fmt"

func main() {
	fmt.Println("1. Bentuk for lengkap:")
	// Cetak angka 0 sampai 4
	for i := 0; i < 5; i++ {
		fmt.Printf("Iterasi ke-%d\n", i)
	}

	fmt.Println("======================================")
	fmt.Println("2. Bentuk for seperti 'while':")
	n := 0
	for n < 3 {
		fmt.Println("Nilai n:", n)
		n++
	}

	fmt.Println("======================================")
	fmt.Println("3. Contoh 'continue' dan 'break':")
	// Cetak hanya angka ganjil dari 0-9, tapi berhenti jika sudah mencapai 8
	for i := 0; i < 10; i++ {
		if i == 8 {
			fmt.Println("Loop dihentikan oleh break!")
			break // Keluar dari loop
		}
		if i%2 == 0 { // Jika i adalah angka genap
			continue // Lewatkan sisa kode di iterasi ini, lanjut ke i berikutnya
		}
		fmt.Printf("Angka ganjil: %d\n", i)
	}

	fmt.Println("======================================")
	fmt.Println("4. Contoh hitung mundur 5 ke 1")
	for i := 5; i > 0; i-- {
		fmt.Printf("Angka: %d\n", i)
		if i == 1 {
			fmt.Println("Meluncur!")
			break
		}
	}
}

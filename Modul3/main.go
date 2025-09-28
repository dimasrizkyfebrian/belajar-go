package main

import "fmt"

func main() {
	// Membuat sebuah slice of string
	buah := []string{"Apel", "Pisang", "Jeruk"}

	fmt.Println("Slice awal:", buah)
	fmt.Printf("Jumlah buah (len): %d\n", len(buah))
	fmt.Printf("Kapasitas (cap): %d\n\n", cap(buah))

	// Menambahkan elemen baru ke slice menggunakan append
	fmt.Println("Menambahkan 'Mangga'...")
	buah = append(buah, "Mangga")

	fmt.Println("Slice setelah ditambah:", buah)
	fmt.Printf("Jumlah buah (len) sekarang: %d\n", len(buah))
	fmt.Printf("Kapasitas (cap) sekarang: %d\n\n", cap(buah))

	// Menambahkan nanas dan anggur
	fmt.Println("Menambahkan 'Nanas dan Anggur'...")
	buah = append(buah, "Nanas", "Anggur")

	fmt.Println("Slice setelah ditambah:", buah)
	fmt.Printf("Jumlah buah (len) sekarang: %d\n", len(buah))
	fmt.Printf("Kapasitas (cap) sekarang: %d\n", cap(buah))
}

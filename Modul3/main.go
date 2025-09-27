package main

import "fmt"

func main() {
	// Membuat array string berisi 4 nama.
	// Go akan otomatis menghitung ukurannya jika kita menggunakan [...]
	namaSiswa := [...]string{"Adi", "Budi", "Cici", "Dodi"}
	namaSiswa[2] = "Citra" // Ubah data array setelah deklarasi

	fmt.Println("Daftar Nama Siswa:")
	fmt.Println(namaSiswa) // Mencetak seluruh array

	// Mengakses dan mencetak elemen satu per satu
	fmt.Println("Siswa pertama:", namaSiswa[0])
	fmt.Println("Siswa ketiga:", namaSiswa[2])

	// Mendapatkan panjang array menggunakan fungsi len()
	jumlahSiswa := len(namaSiswa)
	fmt.Printf("Total siswa: %d orang\n", jumlahSiswa)
}

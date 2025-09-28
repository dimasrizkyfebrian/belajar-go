package main

import "fmt"

func main() {
	// Membuat map kosong
	nilaiSiswa := make(map[string]int)

	// Menambahkan beberapa data
	nilaiSiswa["Adi"] = 85
	nilaiSiswa["Budi"] = 90
	nilaiSiswa["Cici"] = 78

	fmt.Println("Peta nilai siswa:", nilaiSiswa)

	// Mengakses nilai Budi
	fmt.Println("Nilai Budi:", nilaiSiswa["Budi"])

	// Menghapus data Cici
	delete(nilaiSiswa, "Cici")
	fmt.Println("Setelah Cici dihapus:", nilaiSiswa)

	fmt.Println("\n======================================")
	fmt.Println("Mengecek keberadaan key:")

	// Mengecek nilai untuk "Dodi" (yang tidak ada)
	nilaiDodi, ok := nilaiSiswa["Dodi"]
	if ok {
		fmt.Println("Nilai Dodi adalah:", nilaiDodi)
	} else {
		fmt.Println("Data nilai untuk Dodi tidak ditemukan.")
	}

	fmt.Println("\n======================================")
	fmt.Println("Menambahkan data baru:")

	nilaiSiswa["Dimas"] = 95
	nilaiSiswa["Adi"] = 100

	fmt.Println("Peta nilai siswa:", nilaiSiswa)
}

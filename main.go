package main

import "fmt"

func main() {
	const namaAplikasi = "Aplikasi Belajar Go"
	const versi = "1.0"
	const asalNegara = "Indonesia"

	var umur int = 22
	var tinggiBadan float64 = 172.5
	var nomorAbsen uint = 10

	fmt.Println("Selamat datang di", namaAplikasi, "versi", versi)
	fmt.Println("===============================================")

	fmt.Printf("Umur saya: %d tahun\n", umur)
	fmt.Printf("Nomor absen saya adalah: %d\n", nomorAbsen)
	fmt.Printf("Tinggi badan saya adalah %.1f cm\n", tinggiBadan)
	fmt.Printf("Saya berasal dari negara %s\n", asalNegara)
}

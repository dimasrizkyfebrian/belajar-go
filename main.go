package main

import "fmt"

func main() {
	namaLengkap := "Dimas Rizky Febrian"
	pekerjaan := "Calon Go Developer"
	umur := 22
	kotaAsal := "Bekasi"
	apakahSudahMakan := true

	fmt.Println("Halo, nama lengkap saya:")
	fmt.Println(namaLengkap)

	fmt.Println("Pekerjaan saya adalah:")
	fmt.Println(pekerjaan)

	fmt.Println("Saya berasal dari kota", kotaAsal)
	fmt.Println("Umur saya:", umur, "tahun")
	fmt.Println("Sudah makan?", apakahSudahMakan)
}

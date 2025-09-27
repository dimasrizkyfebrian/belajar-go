package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat1 := "Halo, nama saya"
	nama := "Dimas"

	kalimatLengkap := kalimat1 + " " + nama
	fmt.Println(kalimatLengkap)

	fmt.Printf("Panjang kalimat di atas adalah: %d karakter\n", len(kalimatLengkap))

	fmt.Println("=====================================")
	fmt.Println("     Menggunakan Package Strings     ")
	fmt.Println("=====================================")

	namaBesar := strings.ToUpper(nama)
	fmt.Printf("Nama dalam huruf besar adalah: %s\n", namaBesar)

	namaKecil := strings.ToLower(namaBesar)
	fmt.Printf("Nama dalam huruf kecil adalah: %s\n", namaKecil)

	apakahAdaKataHalo := strings.Contains(kalimatLengkap, "Halo")
	fmt.Printf("Apakah ada kalimat mengandung kata 'Halo'? %t\n", apakahAdaKataHalo)

	apakahAdaKataDunia := strings.Contains(kalimatLengkap, "Dunia")
	fmt.Printf("Apakah ada kalimat mengandung kata 'Dunia'? %t\n", apakahAdaKataDunia)

	kalimatGanti := strings.ReplaceAll(kalimatLengkap, "nama", "panggilan")
	fmt.Println(kalimatGanti)
}

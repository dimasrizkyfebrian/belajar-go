package main

import "fmt"

func main() {
	// Slice awal kita
	angka := []int{0, 2, 4, 6, 8, 10}
	fmt.Println("Slice awal:", angka)

	// Mengambil irisan dari indeks 1 sampai 4 (indeks 4 tidak termasuk)
	// Hasilnya adalah elemen di indeks 1, 2, 3 -> {2, 4, 6}
	sliceTengah := angka[1:4]
	fmt.Println("Slice tengah [1:4]:", sliceTengah)

	// Eksperimen Slice
	sliceEksperimen := angka[2:4]
	fmt.Println("Slice eksperimen [2:4]:", sliceEksperimen)

	// Mengambil irisan dari awal sampai indeks 3 (indeks 3 tidak termasuk)
	sliceAwal := angka[:3]
	fmt.Println("Slice awal [:3]:", sliceAwal)

	// Mengambil irisan dari indeks 3 sampai akhir
	sliceAkhir := angka[3:]
	fmt.Println("Slice akhir [3:]:", sliceAkhir)

	fmt.Println("\n======================================")
	fmt.Println("Membuktikan slice berbagi array yang sama:")

	fmt.Println("Slice akhir sebelum diubah:", sliceAkhir)
	// Kita ubah elemen pertama di sliceAkhir (yaitu angka 6)
	sliceAkhir[0] = 999
	fmt.Println("Slice akhir setelah diubah:", sliceAkhir)

	// Lihat apa yang terjadi pada slice 'angka' yang asli
	fmt.Println("Slice 'angka' asli setelah slice akhir diubah:", angka)
}

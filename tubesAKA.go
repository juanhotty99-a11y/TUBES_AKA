package main

import (
	"fmt"
	"time"
)

const MAX = 100000

type Pemain struct {
	ID   int
	Nama string
	Klub string
}

func BinarySearchIterative(data *[MAX]Pemain, n int, targetID int) int {
	kiri := 0
	kanan := n - 1

	for kiri <= kanan {
		tengah := kiri + (kanan-kiri)/2

		if data[tengah].ID == targetID {
			return tengah
		}
		if data[tengah].ID < targetID {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func BinarySearchRecursive(data *[MAX]Pemain, targetID int, kiri int, kanan int) int {
	if kiri > kanan {
		return -1
	}

	tengah := kiri + (kanan-kiri)/2

	if data[tengah].ID == targetID {
		return tengah
	}

	if data[tengah].ID < targetID {
		return BinarySearchRecursive(data, targetID, tengah+1, kanan)
	}
	return BinarySearchRecursive(data, targetID, kiri, tengah-1)
}

func BuatData(hasil *[MAX]Pemain, jumlah int) {
	var daftarNama [5]string
	daftarNama[0] = "Lionel Messi"
	daftarNama[1] = "Cristiano Ronaldo"
	daftarNama[2] = "Mbappe"
	daftarNama[3] = "Haaland"
	daftarNama[4] = "Neymar"

	var daftarKlub [5]string
	daftarKlub[0] = "Inter Miami"
	daftarKlub[1] = "Al Nassr"
	daftarKlub[2] = "Real Madrid"
	daftarKlub[3] = "Man City"
	daftarKlub[4] = "Al Hilal"

	i := 0
	for i < jumlah {
		hasil[i].ID = (i + 1) * 10
		idx := i % 5
		hasil[i].Nama = daftarNama[idx]
		hasil[i].Klub = daftarKlub[idx]
		i++
	}
}

func MulaiBenchmark() {
	var ukuran [5]int
	ukuran[0] = 10
	ukuran[1] = 100
	ukuran[2] = 1000
	ukuran[3] = 10000
	ukuran[4] = 100000

	pengulangan := 200000

	fmt.Println("\n=== HASIL BENCHMARK (Rata-rata Nanodetik per 1x pencarian) ===")
	fmt.Println("Pengulangan per skenario =", pengulangan)
	fmt.Printf("%-10s | %-10s | %-15s | %-15s\n", "N", "Skenario", "Iteratif (ns)", "Rekursif (ns)")
	fmt.Println("---------------------------------------------------------------")

	var data [MAX]Pemain

	i := 0
	for i < 5 {
		n := ukuran[i]
		BuatData(&data, n)

		targetAwal := data[0].ID
		targetTengah := data[n/2].ID
		targetAkhir := data[n-1].ID
		targetTidakAda := data[n-1].ID + 7

		BenchmarkSkenario(&data, n, targetAwal, "Awal", pengulangan)
		BenchmarkSkenario(&data, n, targetTengah, "Tengah", pengulangan)
		BenchmarkSkenario(&data, n, targetAkhir, "Akhir", pengulangan)
		BenchmarkSkenario(&data, n, targetTidakAda, "TidakAda", pengulangan)

		fmt.Println("---------------------------------------------------------------")
		i++
	}

	fmt.Println("*Catatan: biasanya iteratif sedikit lebih cepat karena tidak ada overhead rekursi.")
}

func BenchmarkSkenario(data *[MAX]Pemain, n int, target int, label string, pengulangan int) {

	dummy := 0

	start1 := time.Now()
	k := 0
	for k < pengulangan {
		dummy = dummy + BinarySearchIterative(data, n, target)
		k++
	}
	total1 := time.Since(start1).Nanoseconds()
	rata1 := total1 / int64(pengulangan)

	start2 := time.Now()
	k = 0
	for k < pengulangan {
		dummy = dummy + BinarySearchRecursive(data, target, 0, n-1)
		k++
	}
	total2 := time.Since(start2).Nanoseconds()
	rata2 := total2 / int64(pengulangan)

	if dummy == -999999 {
		fmt.Println("dummy")
	}

	fmt.Printf("%-10d | %-10s | %-15d | %-15d\n", n, label, rata1, rata2)
}

func main() {
	var dataDemo [MAX]Pemain
	nDemo := 20
	BuatData(&dataDemo, nDemo)

	for {
		fmt.Println("\n=== APLIKASI DATA PEMAIN BOLA ===")
		fmt.Println("1. Lihat Data Pemain")
		fmt.Println("2. Cari Pemain (Iteratif)")
		fmt.Println("3. Cari Pemain (Rekursif)")
		fmt.Println("4. Cek Kecepatan (Benchmark)")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			i := 0
			for i < nDemo {
				p := dataDemo[i]
				fmt.Printf("ID: %d | %s (%s)\n", p.ID, p.Nama, p.Klub)
				i++
			}

		} else if pilihan == 2 {
			fmt.Print("Masukkan ID (kelipatan 10, contoh: 10, 20, 30): ")
			var cari int
			fmt.Scan(&cari)

			idx := BinarySearchIterative(&dataDemo, nDemo, cari)
			if idx != -1 {
				p := dataDemo[idx]
				fmt.Printf("Ketemu (Iteratif): %s (%s)\n", p.Nama, p.Klub)
			} else {
				fmt.Println("Tidak ditemukan")
			}

		} else if pilihan == 3 {
			fmt.Print("Masukkan ID (kelipatan 10, contoh: 10, 20, 30): ")
			var cari int
			fmt.Scan(&cari)

			idx := BinarySearchRecursive(&dataDemo, cari, 0, nDemo-1)
			if idx != -1 {
				p := dataDemo[idx]
				fmt.Printf("Ketemu (Rekursif): %s (%s)\n", p.Nama, p.Klub)
			} else {
				fmt.Println("Tidak ditemukan")
			}

		} else if pilihan == 4 {
			MulaiBenchmark()

		} else if pilihan == 5 {
			return

		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

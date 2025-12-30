package main

import (
	"fmt"
	"time"
)

// 1. WADAH DATA
type Pemain struct {
	ID   int
	Nama string
	Klub string
}

// 2. ALGORITMA ITERATIF
func BinarySearchIterative(data []Pemain, targetID int) int {
	kiri := 0
	kanan := len(data) - 1

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

// 3. ALGORITMA REKURSIF
func BinarySearchRecursive(data []Pemain, targetID int, kiri int, kanan int) int {
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

// 4. PEMBUAT DATA DUMMY
func BuatData(jumlah int) []Pemain {
	daftarNama := []string{"Lionel Messi", "Cristiano Ronaldo", "Mbappe", "Haaland", "Neymar"}
	daftarKlub := []string{"Inter Miami", "Al Nassr", "Real Madrid", "Man City", "Al Hilal"}

	var hasil []Pemain

	// Loop biasa untuk isi data
	for i := 0; i < jumlah; i++ {
		var p Pemain
		p.ID = (i + 1) * 10

		// Ambil nama gantian (pakai sisa bagi/modulo)
		idx := i % 5
		p.Nama = daftarNama[idx]
		p.Klub = daftarKlub[idx]

		hasil = append(hasil, p)
	}
	return hasil
}

// 5. FITUR BENCHMARK (INI YANG DIUBAH JADI LOOP BIASA)
// 5. FITUR BENCHMARK (DIPERBAIKI BIAR GAK 0)
func MulaiBenchmark() {
	ukuran := []int{10, 100, 1000, 10000, 100000}

	// Kita akan mengulang pencarian sebanyak 200.000 kali
	// supaya waktunya bisa terukur (tidak 0 lagi)
	pengulangan := 200000

	fmt.Println("\n=== HASIL CEK KECEPATAN (Rata-rata Nanodetik) ===")
	fmt.Printf("%-10s | %-15s | %-15s\n", "Jumlah Data", "Iteratif (ns)", "Rekursif (ns)")
	fmt.Println("------------------------------------------------")

	for i := 0; i < 5; i++ {
		n := ukuran[i]

		// 1. Buat Data
		data := BuatData(n)
		target := data[n-1].ID // Cari data paling ujung (paling susah)

		// --- CEK ITERATIF ---
		start1 := time.Now()
		// Ulangi 200.000 kali biar kerasa berat
		for k := 0; k < pengulangan; k++ {
			BinarySearchIterative(data, target)
		}
		totalWaktu1 := time.Since(start1).Nanoseconds()
		rataRata1 := totalWaktu1 / int64(pengulangan) // Hitung rata-rata per 1x cari

		// --- CEK REKURSIF ---
		start2 := time.Now()
		// Ulangi 200.000 kali juga
		for k := 0; k < pengulangan; k++ {
			BinarySearchRecursive(data, target, 0, len(data)-1)
		}
		totalWaktu2 := time.Since(start2).Nanoseconds()
		rataRata2 := totalWaktu2 / int64(pengulangan) // Hitung rata-rata

		// Cetak Hasil
		fmt.Printf("%-10d | %-15d | %-15d\n", n, rataRata1, rataRata2)
	}

	fmt.Println("------------------------------------------------")
	fmt.Println("*Hasil adalah rata-rata waktu untuk 1x pencarian.")
}

// 6. MAIN MENU
func main() {
	dataDemo := BuatData(20)

	for {
		fmt.Println("\n=== APLIKASI DATA PEMAIN BOLA ===")
		fmt.Println("1. Lihat Data Pemain")
		fmt.Println("2. Cari Pemain")
		fmt.Println("3. Cek Kecepatan (Benchmark)")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			for i := 0; i < len(dataDemo); i++ {
				p := dataDemo[i]
				fmt.Printf("ID: %d | %s (%s)\n", p.ID, p.Nama, p.Klub)
			}
		} else if pilihan == 2 {
			fmt.Print("Masukkan ID: ")
			var cari int
			fmt.Scan(&cari)

			idx := BinarySearchIterative(dataDemo, cari)
			if idx != -1 {
				p := dataDemo[idx]
				fmt.Printf("Ketemu: %s (%s)\n", p.Nama, p.Klub)
			} else {
				fmt.Println("Tidak ditemukan")
			}
		} else if pilihan == 3 {
			MulaiBenchmark()
		} else if pilihan == 4 {
			return
		}
	}
}

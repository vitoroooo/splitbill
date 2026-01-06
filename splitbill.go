package main

import "fmt"

type Orang struct {
	nama   string
	bayar  float64
	hutang float64
}

var anggota []Orang
var totalBiaya float64

func main() {
	fmt.Println("\n╔════════════════════════════════╗")
	fmt.Println("║  KALKULATOR SPLIT BILL          ║")
	fmt.Println("╚════════════════════════════════╝\n")

	if !inputAnggota() {
		return
	}

	if !inputPengeluaran() {
		return
	}

	hitungPembagian()
	tampilkanHasil()
}

func inputAnggota() bool {
	var jumlah int
	fmt.Print("Jumlah anggota: ")
	if _, err := fmt.Scan(&jumlah); err != nil || jumlah <= 0 {
		fmt.Println("❌ Jumlah anggota tidak valid.")
		return false
	}

	for i := 0; i < jumlah; i++ {
		var nama string
		fmt.Printf("Nama anggota %d: ", i+1)
		fmt.Scan(&nama)
		anggota = append(anggota, Orang{nama: nama})
	}
	fmt.Println()
	return true
}

func inputPengeluaran() bool {
	var jumlahTransaksi int
	fmt.Print("Jumlah transaksi: ")
	if _, err := fmt.Scan(&jumlahTransaksi); err != nil || jumlahTransaksi <= 0 {
		fmt.Println("❌ Jumlah transaksi tidak valid.")
		return false
	}
	fmt.Println()

	for i := 0; i < jumlahTransaksi; i++ {
		var nama string
		var jumlah float64

		fmt.Printf("Transaksi %d\n", i+1)
		fmt.Print("  Siapa yang bayar: ")
		fmt.Scan(&nama)

		fmt.Print("  Jumlah: Rp ")
		if _, err := fmt.Scan(&jumlah); err != nil || jumlah <= 0 {
			fmt.Println("❌ Jumlah uang tidak valid.")
			return false
		}

		ditemukan := false
		for j := 0; j < len(anggota); j++ {
			if anggota[j].nama == nama {
				anggota[j].bayar += jumlah
				totalBiaya += jumlah
				ditemukan = true
				break
			}
		}

		if !ditemukan {
			fmt.Println("❌ Nama tidak terdaftar. Transaksi dibatalkan.")
			return false
		}

		fmt.Println()
	}
	return true
}

func hitungPembagian() {
	bagianPerOrang := totalBiaya / float64(len(anggota))
	for i := 0; i < len(anggota); i++ {
		anggota[i].hutang = bagianPerOrang - anggota[i].bayar
	}
}

func tampilkanHasil() {
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("              RINGKASAN")
	fmt.Println("═══════════════════════════════════════")
	fmt.Printf("Total Pengeluaran: Rp %.0f\n", totalBiaya)
	fmt.Printf("Bagian per orang: Rp %.0f\n\n", totalBiaya/float64(len(anggota)))

	for _, a := range anggota {
		fmt.Printf("%s\n", a.nama)
		fmt.Printf("  Bayar: Rp %.0f\n", a.bayar)

		if a.hutang > 0 {
			fmt.Printf("  Harus bayar: Rp %.0f\n\n", a.hutang)
		} else if a.hutang < 0 {
			fmt.Printf("  Harus terima: Rp %.0f\n\n", -a.hutang)
		} else {
			fmt.Println("  Lunas\n")
		}
	}
}

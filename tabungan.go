package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func input(msg string) string {
	fmt.Print(msg)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func daftarMenu() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Tabungan Golang")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("1. Tampilkan daftar tabungan")
	fmt.Println("2. Buat daftar tabungaan")
	fmt.Println("3. Isi tabungan")
	fmt.Println("4. Ambil isi tabungan")
	fmt.Println("5. Edit tabungan")
	fmt.Println("6. Hapus tabungan")
	fmt.Println("0. Keluar")
}

type tabungan struct {
	nama      string
	target    float64
	terkumpul float64
	status    string
}

func buatTabunganBaru(daftar []tabungan, namaTabungan string, targetTabungan float64) []tabungan {
	tabunganBaru := tabungan{nama: namaTabungan, target: targetTabungan, terkumpul: 0.0, status: "proses"}
	fmt.Printf("\nTabungan %s berhasil dibuat\n\n", tabunganBaru.nama)
	return append(daftar, tabunganBaru)
}

func perbaruiStatus(t *tabungan) {
    if t.terkumpul >= t.target {
        t.status = "terpenuhi"
    } else {
        t.status = "proses"
    }
}

func tampilkanTabungan (daftar []tabungan) {
	if len(daftar) == 0 {
		fmt.Printf("Tidak ada tabungan\n\n")
	} else {
		fmt.Println()
		fmt.Println(strings.Repeat("=", 100))
		fmt.Printf("%-4s %-25s %-18s %-18s %-18s %-10s\n", "No", "Nama", "Target", "Terkumpul", "Kekurangan", "Status")

		fmt.Println(strings.Repeat("-", 100))
		for i, v := range daftar {
			fmt.Printf("%-4d %-25sRp. %-18.0fRp. %-18.0fRp. %-18.0f %-10s\n", i+1, v.nama, v.target, v.terkumpul, (v.target - v.terkumpul), v.status)
		}
		fmt.Println(strings.Repeat("=", 100))
		fmt.Println()
	}
}

func isiTabungan(daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if (strings.EqualFold(v.nama, namaTabungan)){
			nominal, err := strconv.ParseFloat(input("Masukan nominal: Rp "), 64)
			if err != nil {
				fmt.Println("Nominal harus berupa angka!")
				return
			}
			daftar[i].terkumpul += nominal
			perbaruiStatus(&daftar[i])
			fmt.Printf("\nTabungan %q berhasil diisi\n\n", v.nama)
			return
		}
	}
	fmt.Printf("\nNama tabungan tidak ditemukan\n\n")
}

func ambilIsiTabungan (daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if (strings.EqualFold(v.nama, namaTabungan)){
			nominal, err := strconv.ParseFloat(input("Nomimal Pengambilan : Rp "), 64)
			if err != nil {
				fmt.Println("Nominal harus berupa angka!")
				return
			}
			if nominal > daftar[i].terkumpul {
				fmt.Println("Nominal tidak mencukupi!")
				return
			}
			daftar[i].terkumpul -= nominal
			perbaruiStatus(&daftar[i])
			fmt.Printf("\nTabungan %q berhasil diambil\n\n", v.nama)
			return
		}
	}
	fmt.Printf("\nNama tabungan tidak ditemukan!\n\n")
}

func editTabungan (daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if (strings.EqualFold(v.nama, namaTabungan)){
			namaBaru := input("Masukan nama baru: ")
			targetBaru, err := strconv.ParseFloat(input("Masukan target baru: Rp "), 64)
			if err != nil {
				fmt.Println("Nominal harus berupa angka!")
				return
			}
			daftar[i].nama = namaBaru
			daftar[i].target = targetBaru
			fmt.Printf("\nTabungan Berhasil diedit\n\n")
			return
		}
	}
	fmt.Printf("\nNama tabungan tidak ditemukan!\n\n")
}

func hapusTabungan (daftar []tabungan, namaTabungan string) []tabungan {
	for i, v := range daftar {
		if (strings.EqualFold(v.nama, namaTabungan)){
			fmt.Printf("\nTabungan %q berhasil dihapus\n\n", v.nama)
			return append(daftar[:i], daftar[i+1:]...)
		}
	}
	fmt.Printf("\nNama tabungan tidak ditemukan!\n\n")
	return daftar
}

func main() {
	daftarTabungan := []tabungan{}

	for {
		daftarMenu()
		
		fmt.Println()
		fmt.Println(strings.Repeat("-", 50))
		pilihan, _ := strconv.Atoi(input("Masukan pilihan: "))

		if pilihan == 0 {
			break
		}

		switch pilihan {
		case 1:
			tampilkanTabungan(daftarTabungan)
		case 2:
			nama := input("Masukan nama tabungan: ")
			target, err := strconv.ParseFloat(input("Masukan target tabungan: Rp "), 64)

			if err != nil {
				fmt.Printf("\nMasukan harus berupa angka!\n")
			} else {
				daftarTabungan = buatTabunganBaru(daftarTabungan, nama, target)
			}
		case 3:
			nama := input("Masukan nama tabungan: ")
			isiTabungan(daftarTabungan, nama)
		case 4:
			nama := input("Masukan nama tabungan: ")
			ambilIsiTabungan(daftarTabungan, nama)
		case 5:
			nama := input("Masukan nama tabungan yang ingin diedit: ")
			editTabungan(daftarTabungan, nama)
		case 6:
			nama := input("Masukan nama tabungan yang ingin dihapus: ")
			daftarTabungan = hapusTabungan(daftarTabungan, nama)
		}
	}

}

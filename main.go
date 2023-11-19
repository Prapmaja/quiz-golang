package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struktur untuk menyimpan informasi pengguna dan hasil kuis
type HasilKuis struct {
	Nama         string
	Skor         int
	JawabanBenar int
	JawabanSalah int
}

// Struktur untuk menyimpan pertanyaan dan pilihan jawaban
type Pertanyaan struct {
	Pertanyaan   string
	Pilihan      []string
	JawabanBenar int
}

func main() {
	// Membuat array pertanyaan
	pertanyaan := []Pertanyaan{
		{
			Pertanyaan:   "Kota apakah yang berawalan huruf D?",
			Pilihan:      []string{"Denpasar", "Jogja", "Bandung", "Tangerang"},
			JawabanBenar: 0,
		},
		{
			Pertanyaan:   "Hewan apakah yang bisa terbang?",
			Pilihan:      []string{"Agus", "Kucing", "Elang", "Tikus"},
			JawabanBenar: 2,
		},
		// Tambahkan pertanyaan lain di sini jika diperlukan
	}

	// Membuat variabel hasil kuis
	var hasilKuis HasilKuis

	// Meminta nama pengguna
	fmt.Print("Input nama: ")
	hasilKuis.Nama = getInput()

	// Menampilkan pertanyaan
	for _, pertanyaan := range pertanyaan {
		// Menampilkan pertanyaan
		fmt.Println(pertanyaan.Pertanyaan)

		// Menampilkan pilihan jawaban
		for j, pilihan := range pertanyaan.Pilihan {
			fmt.Printf("%d. %s\n", j, pilihan)
		}

		// Meminta jawaban dari pengguna
		fmt.Print("Masukkan jawaban (0,1,2,3): ")
		jawaban := getJawaban()

		// Memeriksa jawaban
		if jawaban == pertanyaan.JawabanBenar {
			fmt.Println()
			hasilKuis.JawabanBenar++
			hasilKuis.Skor += 1
		} else {
			fmt.Println()
			hasilKuis.JawabanSalah++
		}
	}

	// Menampilkan hasil kuis
	fmt.Printf("\nStatistic Kuis\n")
	fmt.Printf("Nama: %s\n", hasilKuis.Nama)
	fmt.Printf("Skor: %d\n", hasilKuis.Skor)
	fmt.Printf("Jawaban Benar: %d\n", hasilKuis.JawabanBenar)
	fmt.Printf("Jawaban Salah: %d\n", hasilKuis.JawabanSalah)
}

// Fungsi untuk membaca input dari pengguna
func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Fungsi untuk membaca input jawaban dari pengguna
func getJawaban() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Konversi input ke tipe data int
	jawaban, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Masukkan angka yang valid.")
		return getJawaban()
	}

	return jawaban
}

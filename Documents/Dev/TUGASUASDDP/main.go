//KELOMPOK 3

// Agus Putra Prapmaja - 2301020051
// I Kadek Momet Dwika Putra - 2301020037
// Putu Yoga Aditya Perdana Putra - 2301020034

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Karyawan adalah struktur data untuk merepresentasikan informasi karyawan.
type Karyawan struct {
	Nama                string
	PunyaIstri          bool
	PunyaAnak           bool
	JumlahAnak          int
	JumlahPengantaran   int
	PengantaranBerhasil int
}

//menghitung gaji karyawan berdasarkan aturan tertentu.
func (k *Karyawan) HitungGaji() int {
	const gajiPokok = 4000000
	const tunjanganIstri = 1000000
	const upahPengantaran = 10000

	gaji := gajiPokok

	if k.PunyaIstri {
		gaji += tunjanganIstri
	}
	gaji += k.JumlahAnak * 500000
	gaji += k.JumlahPengantaran * upahPengantaran

	if k.PengantaranBerhasil > 50 {
		gaji += 1000000
	} else if k.PengantaranBerhasil > 10 {
		gaji += 500000
	}

	if k.PengantaranBerhasil < 5 {
		pemotongan := 500000
		gaji -= pemotongan
	}

	return gaji
}

//menambahkan informasi pengantaran untuk seorang karyawan
func (k *Karyawan) TambahPengantaran(berhasil bool) {
	k.JumlahPengantaran++
	if berhasil {
		k.PengantaranBerhasil++
	}
}

func (k *Karyawan) TampilkanInfo() string {
	return fmt.Sprintf("Nama Karyawan: %s\nKeterangan: %s dan memiliki %d anak\n", k.Nama, func() string {
		if k.PunyaIstri {
			return "Sudah Menikah"
		}
		return "Belum Menikah"
	}(), k.JumlahAnak)
}

// Menampilkan Sebuah menu setelah dirun
func main() {
	for {
		fmt.Println("Program Penggajian Karyawan")
		fmt.Println("==========================")
		fmt.Println("1. Tambah Karyawan")
		fmt.Println("2. Cari Karyawan")
		fmt.Println("3. Hapus Karyawan")
		fmt.Println("4. Tambah Pengantaran")
		fmt.Println("5. Cari Pengantaran")
		fmt.Println("6. Hapus Pengantaran")
		fmt.Println("7. Hitung Gaji")
		fmt.Println("8. Keluar")

		//kode untuk memilih pilihan yang tersedia sesuai fungsinya
		var pilihan int
		fmt.Print("Pilih Menu [1-8]: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			TambahKaryawan()
		case 2:
			CariKaryawan()
		case 3:
			HapusKaryawan()
		case 4:
			TambahPengantaran()
		case 5:
			CariPengantaran()
		case 6:
			HapusPengantaran()
		case 7:
			HitungGaji()
		case 8:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

//meminta input dari pengguna dan menambahkan karyawan baru.
func TambahKaryawan() {
	var nama string
	var punyaIstri, punyaAnak string

	// Meminta input nama karyawan dari pengguna
	fmt.Print("Nama Karyawan: ")
	fmt.Scan(&nama)

	// Meminta input status memiliki istri (y/n) dari pengguna
	fmt.Print("Punya Istri (y/n): ")
	fmt.Scan(&punyaIstri)

	var jumlahAnak int
	if punyaIstri == "y" {
		// Jika memiliki istri, meminta input status memiliki anak (y/n) dari pengguna.
		fmt.Print("Punya Anak (y/n): ")
		fmt.Scan(&punyaAnak)

		// Jika memiliki anak, meminta input jumlah anak dari pengguna.
		if punyaAnak == "y" {
			fmt.Print("Masukkan jumlah anak: ")
			fmt.Scan(&jumlahAnak)
		}
	}

	// Membuat objek Karyawan baru berdasarkan input pengguna.
	karyawanBaru := Karyawan{
		Nama:       nama,
		PunyaIstri: punyaIstri == "y",
		PunyaAnak:  punyaAnak == "y",
		JumlahAnak: jumlahAnak,
	}

	// Menyimpan data Karyawan baru ke dalam file JSON.
	saveData(karyawanBaru)

	// Menampilkan pesan sukses.
	fmt.Println("Karyawan berhasil ditambahkan!")
}

//meminta nama karyawan dari pengguna dan menampilkan informasi karyawan jika ditemukan.
func CariKaryawan() {
	// Memuat data karyawan dari file JSON.
	saveDatas := loadData()

	// Meminta input nama karyawan yang akan dicari.
	var namaCari string
	fmt.Print("Masukkan nama karyawan: ")
	fmt.Scan(&namaCari)

	// Variabel untuk menyimpan hasil pencarian karyawan.
	var karyawan *Karyawan

	// Iterasi melalui daftar karyawan untuk mencari karyawan dengan nama yang sesuai.
	for i := range saveDatas {
		if saveDatas[i].Nama == namaCari {
			karyawan = &saveDatas[i]
			break
		}
	}

	// Menampilkan hasil pencarian.
	if karyawan != nil {
		// Jika karyawan ditemukan, menampilkan informasi karyawan.
		fmt.Print(karyawan.TampilkanInfo())
	} else {
		// Jika karyawan tidak ditemukan, menampilkan pesan bahwa karyawan tidak ditemukan.
		fmt.Println("Karyawan tidak ditemukan!")
	}
}

//menampilkan daftar karyawan, meminta input pengguna untuk memilih karyawan, dan menghapusnya jika valid.
func HapusKaryawan() {
	// Memuat data karyawan dari file JSON.
	saveDatas := loadData()

	// Menampilkan daftar karyawan yang dapat dihapus.
	fmt.Println("Daftar Karyawan:")
	for i := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, saveDatas[i].Nama)
	}

	// Meminta input dari pengguna untuk memilih karyawan yang akan dihapus.
	var indexHapus int
	fmt.Print("Pilih Karyawan yang akan dihapus [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexHapus)
	indexHapus--

	// Memvalidasi pilihan pengguna dan menghapus karyawan jika valid.
	if indexHapus >= 0 && indexHapus < len(saveDatas) {
		// Menghapus karyawan dari slice menggunakan teknik slice append.
		saveDatas = append(saveDatas[:indexHapus], saveDatas[indexHapus+1:]...)

		// Menyimpan data yang telah diubah kembali ke dalam file JSON.
		saveDataA(saveDatas)

		// Menampilkan pesan sukses.
		fmt.Println("Karyawan berhasil dihapus!")
	} else {
		// Menampilkan pesan bahwa pilihan tidak valid.
		fmt.Println("Pilihan tidak valid!")
	}
}

//menampilkan daftar karyawan, meminta input pengguna untuk memilih karyawan, dan menambahkan informasi pengantaran
func TambahPengantaran() {
	// Memuat data karyawan dari file JSON.
	saveDatas := loadData()

	// Menampilkan daftar karyawan yang dapat dipilih untuk pengantaran.
	fmt.Println("Daftar Karyawan:")
	for i, karyawan := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, karyawan.Nama)
	}

	// Meminta input dari pengguna untuk memilih karyawan yang akan melakukan pengantaran.
	var indexPilih int
	fmt.Print("Pilih Karyawan yang akan melakukan pengantaran [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexPilih)
	indexPilih--

	// Memvalidasi pilihan pengguna.
	if indexPilih >= 0 && indexPilih < len(saveDatas) {
		// Meminta informasi tambahan untuk pengantaran.
		var tujuan string
		fmt.Print("Masukkan tujuan pengantaran: ")
		fmt.Scan(&tujuan)

		var berhasil string
		fmt.Print("Apakah pengantaran berhasil (y/n): ")
		fmt.Scan(&berhasil)

		// Memperbarui data karyawan terpilih.
		karyawan := &saveDatas[indexPilih]
		karyawan.JumlahPengantaran++

		if berhasil == "y" {
			karyawan.PengantaranBerhasil++
		}

		// Menyimpan data yang telah diubah kembali ke dalam file JSON.
		saveDataA(saveDatas)

		// Menampilkan pesan sukses.
		fmt.Println("Pengantaran berhasil ditambahkan!")
	} else {
		// Menampilkan pesan bahwa pilihan tidak valid.
		fmt.Println("Pilihan tidak valid!")
	}
}

//menampilkan daftar karyawan, meminta input pengguna untuk memilih karyawan, dan menampilkan informasi pengantaran.
func CariPengantaran() {
	// Memuat data karyawan dari file JSON.
	saveDatas := loadData()

	// Menampilkan daftar karyawan untuk memilih karyawan yang akan dicari pengantarannya.
	fmt.Println("Daftar Karyawan:")
	for i := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, saveDatas[i].Nama)
	}

	// Meminta input dari pengguna untuk memilih karyawan.
	var indexCari int
	fmt.Print("Pilih Karyawan yang akan dicari pengantarannya [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexCari)
	indexCari--

	// Memvalidasi pilihan pengguna dan menampilkan informasi pengantaran untuk karyawan yang dipilih.
	if indexCari >= 0 && indexCari < len(saveDatas) {
		karyawan := &(saveDatas)[indexCari]
		fmt.Printf("Nama Karyawan: %s\n", karyawan.Nama)
		fmt.Printf("Jumlah Pengantaran: %d\n", karyawan.JumlahPengantaran)
		fmt.Printf("Jumlah Pengantaran Berhasil: %d\n", karyawan.PengantaranBerhasil)
	} else {
		// Menampilkan pesan bahwa pilihan tidak valid.
		fmt.Println("Pilihan tidak valid!")
	}
}

//menampilkan daftar karyawan, meminta input pengguna untuk memilih karyawan, dan menghapus informasi pengantaran.
func HapusPengantaran() {
	// Memuat data karyawan dari file JSON.
	saveDatas := loadData()

	// Menampilkan daftar karyawan untuk pengantaran.
	fmt.Println("Daftar Karyawan:")
	for i := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, saveDatas[i].Nama)
	}

	// Meminta input pengguna untuk memilih karyawan yang akan dihapus pengantarannya.
	var indexHapusPengantaran int
	fmt.Print("Pilih Karyawan yang akan dihapus pengantarannya [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexHapusPengantaran)
	indexHapusPengantaran--

	// Memvalidasi pilihan pengguna dan menghapus informasi pengantaran jika valid.
	if indexHapusPengantaran >= 0 && indexHapusPengantaran < len(saveDatas) {
		// Menghapus informasi pengantaran dari karyawan yang dipilih.
		karyawan := &saveDatas[indexHapusPengantaran]
		karyawan.JumlahPengantaran = 0
		karyawan.PengantaranBerhasil = 0

		// Menyimpan data yang telah diubah kembali ke dalam file JSON.
		saveDataA(saveDatas)

		// Menampilkan pesan sukses.
		fmt.Println("Pengantaran berhasil dihapus!")
	} else {
		// Menampilkan pesan bahwa pilihan tidak valid.
		fmt.Println("Pilihan tidak valid!")
	}
}

//menampilkan daftar karyawan, meminta input pengguna untuk memilih karyawan, dan menampilkan gaji karyawan.
func HitungGaji() {
	saveDatas := loadData()
	fmt.Println("Daftar Karyawan:")
	for i := range saveDatas {
		fmt.Printf("%d. %s\n", i+1, saveDatas[i].Nama)
	}

	var indexHitungGaji int
	fmt.Print("Pilih Karyawan yang akan dihitung gajinya [1-", len(saveDatas), "]: ")
	fmt.Scan(&indexHitungGaji)
	indexHitungGaji--

	if indexHitungGaji >= 0 && indexHitungGaji < len(saveDatas) {
		karyawan := saveDatas[indexHitungGaji]
		gaji := karyawan.HitungGaji()
		fmt.Printf("Gaji %s adalah Rp %d\n", karyawan.Nama, gaji)
	} else {
		fmt.Println("Pilihan tidak valid!")
	}
}

// saveDataA menyimpan data karyawan ke file JSON.
func saveDataA(data []Karyawan) {
	// Mendapatkan direktori kerja saat ini.
	wd, _ := os.Getwd()

	// Mengubah data karyawan menjadi format JSON.
	jsonData, jsonError := json.Marshal(data)
	if jsonError != nil {
		log.Fatalln("Can't Marshal the Data")
	}

	// Menyimpan data JSON ke dalam file "save.json".
	writeError := os.WriteFile(fmt.Sprintf("%s/%s", wd, "save.json"), jsonData, os.ModePerm)
	if writeError != nil {
		log.Fatalln("Can't write the file")
	}
}

// saveData menambahkan data karyawan baru ke file JSON atau membuat file baru jika belum ada.
func saveData(data Karyawan) {
	// Slice untuk menyimpan data karyawan sebelum ditambahkan data baru.
	var temps []Karyawan

	// Mendapatkan direktori kerja saat ini.
	wd, _ := os.Getwd()

	// Membaca isi file JSON yang sudah ada.
	oldfile, err := os.ReadFile(fmt.Sprintf("%s/%s", wd, "save.json"))

	// Mengecek apakah file tidak ada (baru).
	if os.IsNotExist(err) {
		// Membuat file baru jika belum ada.
		file, errs := os.Create(fmt.Sprintf("%s/%s", wd, "save.json"))
		if errs != nil {
			log.Fatalln("Error while Creating file")
		}

		// Menambahkan data karyawan baru ke dalam slice.
		temps = append(temps, data)

		// Marshal data ke format JSON.
		jsonData, jsonError := json.Marshal(temps)
		if jsonError != nil {
			log.Fatalln("Failed to Marshal the Data")
		}

		// Menuliskan data JSON ke dalam file.
		_, writeError := file.Write(jsonData)
		if writeError != nil {
			log.Fatalln("Failed to Write the Data")
		}
	}

	// Mengecek apakah ada kesalahan pembacaan file.
	if err != nil {
		log.Fatalln("Failed to Open the file")
	}

	// Mengembalikan data JSON ke dalam slice temp
	jsonErr := json.Unmarshal(oldfile, &temps)
	if jsonErr != nil {
		log.Fatalln("Failed to Unmarshal the data")
	}

	// Menambahkan data karyawan baru ke dalam slice.
	temps = append(temps, data)

	// Marshal data ke format JSON.
	jsonData, jsonError := json.Marshal(temps)
	if jsonError != nil {
		fmt.Println("Failed to Marshal the data")
	}

	// Menuliskan data JSON ke dalam file.
	writeError := os.WriteFile(fmt.Sprintf("%s/%s", wd, "save.json"), jsonData, os.ModePerm)
	if writeError != nil {
		log.Fatalln("Can't write the file")
	}
}

// loadData membaca data karyawan dari file JSON.
func loadData() []Karyawan {
	// Variabel untuk menyimpan data yang akan dibaca dari file JSON.
	var temp []Karyawan

	// Mendapatkan direktori kerja saat ini.
	wd, _ := os.Getwd()

	// Membaca konten file JSON yang berisi data karyawan.
	oldfile, err := os.ReadFile(fmt.Sprintf("%s/%s", wd, "save.json"))

	// Memeriksa apakah file tidak ditemukan.
	if os.IsNotExist(err) {
		log.Fatalln("Can't read the file!, is the file already exists?")
	}

	// Menguraikan konten file JSON ke dalam slice Karyawan.
	jsonError := json.Unmarshal(oldfile, &temp)
	if jsonError != nil {
		log.Fatalln("Can't unmarshal the data")
	}

	// Mengembalikan slice Karyawan yang berisi data yang dibaca.
	return temp
}

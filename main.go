package main

import (
	"fmt"
	"math"
)

// --- Konstanta & tipe bentukan ---
const MAX = 100

const (
	StatusPending  = 0
	StatusBerjalan = 1
	StatusSelesai  = 2
)

type Tanggal struct {
	Hari, Bulan, Tahun int
}

type Proyek struct {
	ID         int
	NamaProyek string
	NamaKlien  string
	Deadline   Tanggal
	Bayaran    int
	Status     int
}

// Variabel global (hanya array utama + jumlah data)
var daftar [MAX]Proyek
var jumlah int

// --- Input & tampilan ---

func bacaKata(label string) string {
	fmt.Print(label + ": ")
	var s string
	fmt.Scan(&s)
	return s
}

func bacaAngka(label string) int {
	fmt.Print(label + ": ")
	var n int
	fmt.Scan(&n)
	return n
}

func teksStatus(s int) string {
	if s == StatusPending {
		return "Pending"
	}
	if s == StatusBerjalan {
		return "Sedang dikerjakan"
	}
	if s == StatusSelesai {
		return "Selesai"
	}
	return "-"
}

func tampilkanTabel() {
	if jumlah == 0 {
		fmt.Println("Data kosong.")
		return
	}
	fmt.Println("\nNo  ID   Proyek         Klien        Deadline    Bayaran     Status")
	for i := 0; i < jumlah; i++ {
		p := daftar[i]
		fmt.Printf("%-3d %-4d %-14s %-12s %02d-%02d-%d %-11d %s\n",
			i+1, p.ID, p.NamaProyek, p.NamaKlien,
			p.Deadline.Hari, p.Deadline.Bulan, p.Deadline.Tahun,
			p.Bayaran, teksStatus(p.Status))
	}
}

func tampilkanDetail(i int) {
	p := daftar[i]
	fmt.Printf("ID:%d | %s | Klien:%s | %02d-%02d-%d | Rp%d | %s\n",
		p.ID, p.NamaProyek, p.NamaKlien,
		p.Deadline.Hari, p.Deadline.Bulan, p.Deadline.Tahun,
		p.Bayaran, teksStatus(p.Status))
}

// --- Bantuan ---

func nilaiTanggal(t Tanggal) int {
	return t.Tahun*10000 + t.Bulan*100 + t.Hari
}

func kecil(s string) string {
	h := ""
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c = c + 32
		}
		h = h + string(c)
	}
	return h
}

func sama(a, b string) bool {
	return kecil(a) == kecil(b)
}

func sebelum(a, b string) bool {
	aa, bb := kecil(a), kecil(b)
	for i := 0; i < len(aa) && i < len(bb); i++ {
		if aa[i] < bb[i] {
			return true
		}
		if aa[i] > bb[i] {
			return false
		}
	}
	return len(aa) < len(bb)
}

func idBaru() int {
	m := 0
	for i := 0; i < jumlah; i++ {
		if daftar[i].ID > m {
			m = daftar[i].ID
		}
	}
	return m + 1
}

func indeksID(id int) int {
	for i := 0; i < jumlah; i++ {
		if daftar[i].ID == id {
			return i
		}
	}
	return -1
}

func tukar(i, j int) {
	temp := daftar[i]
	daftar[i] = daftar[j]
	daftar[j] = temp
}

func isiProyek(p *Proyek) {
	fmt.Println("Isi satu nilai per baris, lalu Enter.")
	p.NamaProyek = bacaKata("Nama proyek")
	p.NamaKlien = bacaKata("Nama klien")
	p.Deadline.Hari = bacaAngka("Hari deadline")
	p.Deadline.Bulan = bacaAngka("Bulan")
	p.Deadline.Tahun = bacaAngka("Tahun")
	p.Bayaran = bacaAngka("Bayaran (Rp)")
	fmt.Println("Status: 0=Pending 1=Berjalan 2=Selesai")
	p.Status = bacaAngka("Status")
}

// --- CRUD ---

func tambah() {
	if jumlah >= MAX {
		fmt.Println("Array penuh.")
		return
	}
	var p Proyek
	p.ID = idBaru()
	isiProyek(&p)
	if p.Status < 0 || p.Status > 2 {
		fmt.Println("Status harus 0, 1, atau 2.")
		return
	}
	daftar[jumlah] = p
	jumlah++
	fmt.Println("Tersimpan.")
}

func ubah(idx int) {
	isiProyek(&daftar[idx])
	fmt.Println("Diubah.")
}

func hapus(idx int) {
	fmt.Print("Hapus? (y/n): ")
	var y string
	fmt.Scan(&y)
	if len(y) == 0 || (y[0] != 'y' && y[0] != 'Y') {
		fmt.Println("Batal.")
		return
	}
	for i := idx; i < jumlah-1; i++ {
		daftar[i] = daftar[i+1]
	}
	jumlah--
	fmt.Println("Terhapus.")
}

func ubahStatus() {
	tampilkanTabel()
	id := bacaAngka("ID proyek")
	idx := indeksID(id)
	if idx < 0 {
		fmt.Println("ID tidak ada.")
		return
	}
	fmt.Println("Status: 0=Pending 1=Berjalan 2=Selesai")
	st := bacaAngka("Status baru")
	if st < 0 || st > 2 {
		fmt.Println("Status tidak valid.")
		return
	}
	daftar[idx].Status = st
	fmt.Println("Status diperbarui.")
}

// --- Pencarian ---

func cariUrut(salinan *[MAX]Proyek, n int, byKlien bool) {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			var a, b string
			if byKlien {
				a, b = salinan[j].NamaKlien, salinan[min].NamaKlien
			} else {
				a, b = salinan[j].NamaProyek, salinan[min].NamaProyek
			}
			if sebelum(a, b) {
				min = j
			}
		}
		if min != i {
			salinan[i], salinan[min] = salinan[min], salinan[i]
		}
	}
}

func cariSequential(kata string, byKlien bool) int {
	for i := 0; i < jumlah; i++ {
		if byKlien {
			if sama(daftar[i].NamaKlien, kata) {
				return i
			}
		} else if sama(daftar[i].NamaProyek, kata) {
			return i
		}
	}
	return -1
}

func cariBinary(kata string, byKlien bool) int {
	var salinan [MAX]Proyek
	for i := 0; i < jumlah; i++ {
		salinan[i] = daftar[i]
	}
	cariUrut(&salinan, jumlah, byKlien)

	kiri, kanan := 0, jumlah-1
	hasil := -1
	ditemukan := false
	for kiri <= kanan && !ditemukan {
		tengah := int(math.Floor(float64(kiri+kanan) / 2.0))
		nama := salinan[tengah].NamaProyek
		if byKlien {
			nama = salinan[tengah].NamaKlien
		}
		if sama(nama, kata) {
			ditemukan = true
			hasil = indeksID(salinan[tengah].ID)
		} else if sebelum(nama, kata) {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return hasil
}

func cari() int {
	if jumlah == 0 {
		fmt.Println("Data kosong.")
		return -1
	}
	fmt.Println("\n1 Sequential - proyek")
	fmt.Println("2 Sequential - klien")
	fmt.Println("3 Binary - proyek")
	fmt.Println("4 Binary - klien")
	metode := bacaAngka("Pilih")
	kata := bacaKata("Kata kunci")

	idx := -1
	if metode == 1 {
		idx = cariSequential(kata, false)
	} else if metode == 2 {
		idx = cariSequential(kata, true)
	} else if metode == 3 {
		idx = cariBinary(kata, false)
	} else if metode == 4 {
		idx = cariBinary(kata, true)
	} else {
		fmt.Println("Pilihan salah.")
		return -1
	}
	if idx < 0 {
		fmt.Println("Tidak ditemukan.")
	} else {
		fmt.Println("Ketemu:")
		tampilkanDetail(idx)
	}
	return idx
}

// --- Pengurutan ---

func selectionSortDeadline(naik bool) {
	for i := 0; i < jumlah-1; i++ {
		best := i
		for j := i + 1; j < jumlah; j++ {
			va := nilaiTanggal(daftar[j].Deadline)
			vb := nilaiTanggal(daftar[best].Deadline)
			if naik && va < vb {
				best = j
			}
			if !naik && va > vb {
				best = j
			}
		}
		if best != i {
			tukar(i, best)
		}
	}
}

func insertionSortBayaran(naik bool) {
	for i := 1; i < jumlah; i++ {
		key := daftar[i]
		j := i - 1
		lanjut := true
		for lanjut && j >= 0 {
			geser := daftar[j].Bayaran > key.Bayaran
			if !naik {
				geser = daftar[j].Bayaran < key.Bayaran
			}
			if geser {
				daftar[j+1] = daftar[j]
				j--
			} else {
				lanjut = false
			}
		}
		daftar[j+1] = key
	}
}

func urutkan() {
	if jumlah == 0 {
		fmt.Println("Data kosong.")
		return
	}
	fmt.Println("\n1 Selection Sort - deadline naik")
	fmt.Println("2 Selection Sort - deadline turun")
	fmt.Println("3 Insertion Sort - bayaran naik")
	fmt.Println("4 Insertion Sort - bayaran turun")
	p := bacaAngka("Pilih")
	if p == 1 {
		selectionSortDeadline(true)
	} else if p == 2 {
		selectionSortDeadline(false)
	} else if p == 3 {
		insertionSortBayaran(true)
	} else if p == 4 {
		insertionSortBayaran(false)
	} else {
		fmt.Println("Pilihan salah.")
		return
	}
	fmt.Println("Sudah diurutkan:")
	tampilkanTabel()
}

// --- Laporan ---

func laporan() {
	fmt.Println("\n=== LAPORAN ===")
	fmt.Println("[Selesai]")
	ada := false
	for i := 0; i < jumlah; i++ {
		if daftar[i].Status == StatusSelesai {
			tampilkanDetail(i)
			ada = true
		}
	}
	if !ada {
		fmt.Println("(kosong)")
	}

	fmt.Println("\n[Berjalan]")
	ada = false
	for i := 0; i < jumlah; i++ {
		if daftar[i].Status == StatusPending || daftar[i].Status == StatusBerjalan {
			tampilkanDetail(i)
			ada = true
		}
	}
	if !ada {
		fmt.Println("(kosong)")
	}
}

// --- Menu ---

func menu() int {
	fmt.Println("\n=== Manajemen Proyek Freelance ===")
	fmt.Println("1 Tambah  2 Ubah  3 Hapus  4 Status")
	fmt.Println("5 Cari    6 Urut  7 Laporan  8 Tampil  0 Keluar")
	return bacaAngka("Menu")
}

func main() {
	fmt.Println("Aplikasi Manajemen Proyek Freelance")
	keluar := false
	for !keluar {
		p := menu()
		if p == 0 {
			keluar = true
			fmt.Println("Keluar.")
		} else if p == 1 {
			tambah()
		} else if p == 2 {
			idx := cari()
			if idx >= 0 {
				ubah(idx)
			}
		} else if p == 3 {
			idx := cari()
			if idx >= 0 {
				hapus(idx)
			}
		} else if p == 4 {
			ubahStatus()
		} else if p == 5 {
			cari()
		} else if p == 6 {
			urutkan()
		} else if p == 7 {
			laporan()
		} else if p == 8 {
			tampilkanTabel()
		} else {
			fmt.Println("Menu 0-8 saja.")
		}
	}
}

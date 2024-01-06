package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// konfigGrade adalah struktur untuk menyimpan konfigurasi sistem penilaian
type konfigGrade struct {
	BatasA   float64
	BatasB   float64
	BatasC   float64
	BatasD   float64
	BatasMin float64
}

// Mahasiswa adalah struktur untuk menyimpan data mahasiswa
type Mahasiswa struct {
	NO       int
	NPM      string
	Nama     string
	UTS      float64
	UAS      float64
	RataRata float64
	Grade    string
}

func Header() {
	fmt.Println("\n=============================================================================")
	fmt.Println("                       Program Data Nilai Mahasiswa                           ")
	fmt.Println("=============================================================================")
	fmt.Println("                       Dibuat Oleh : Yoga Ardiansyah                           ")
	fmt.Println("                                     2IA25                         ")
	fmt.Println("                                     51422643                         ")
	fmt.Println("=============================================================================")
}

// HitungGrade mengembalikan grade berdasarkan nilai akhir dan konfig sistem penilaian
func HitungGrade(nilaiAkhir float64, konfig konfigGrade) string {
	if nilaiAkhir > 100 || nilaiAkhir < konfig.BatasMin {
		return "X" // Grade X untuk nilai di atas 100 atau nilai di bawah batas minimum
	} else if nilaiAkhir >= konfig.BatasA {
		return "A"
	} else if nilaiAkhir >= konfig.BatasB {
		return "B"
	} else if nilaiAkhir >= konfig.BatasC {
		return "C"
	} else if nilaiAkhir >= konfig.BatasD {
		return "D"
	} else {
		return "E"
	}
}

// CekNPMUnik memeriksa apakah NPM sudah ada dalam data mahasiswa
func CekNPMUnik(npm string, dataMahasiswa []Mahasiswa) bool {
	for _, mahasiswa := range dataMahasiswa {
		if mahasiswa.NPM == npm {
			return false // NPM sudah ada, bukan unik
		}
	}
	return true // NPM unik
}

// InputDataMahasiswa mengembalikan slice Mahasiswa berdasarkan jumlah input
func InputDataMahasiswa(jmlMahasiswa int, konfig konfigGrade, dataMahasiswa []Mahasiswa) []Mahasiswa {
	fmt.Println("=============================================================================")
	fmt.Println("                            Input Data Mahasiswa                            ")
	fmt.Println("=============================================================================")

	for i := 0; i < jmlMahasiswa; {
		fmt.Printf("\nInput data mahasiswa ke-%d\n", i+1)

		var mahasiswa Mahasiswa
		for {
			fmt.Print("NPM: ")
			fmt.Scan(&mahasiswa.NPM)

			// Validasi NPM unik
			if CekNPMUnik(mahasiswa.NPM, dataMahasiswa) {
				break
			}
			fmt.Println("Error: NPM sudah ada. Harap masukkan NPM yang unik.")
		}

		fmt.Print("Nama: ")
		fmt.Scan(&mahasiswa.Nama)

		// Validasi input nilai UTS
		for {
			fmt.Print("Nilai UTS: ")
			nilaiUTS, err := strconv.ParseFloat(getUserInput(), 64)
			if err == nil {
				mahasiswa.UTS = nilaiUTS
				break
			}
			fmt.Println("Error: Input yang Anda masukkan bukan angka.")
		}

		// Validasi input nilai UAS
		for {
			fmt.Print("Nilai UAS: ")
			nilaiUAS, err := strconv.ParseFloat(getUserInput(), 64)
			if err == nil {
				mahasiswa.UAS = nilaiUAS
				break
			}
			fmt.Println("Error: Input yang Anda masukkan bukan angka.")
		}

		// Validasi pembagian nol sebelum perhitungan rata-rata
		if mahasiswa.UTS == mahasiswa.UAS {
			mahasiswa.RataRata = mahasiswa.UTS
		} else {
			mahasiswa.RataRata = (mahasiswa.UTS + mahasiswa.UAS) / 2
		}

		mahasiswa.Grade = HitungGrade(mahasiswa.RataRata, konfig)
		mahasiswa.NO = i + 1

		dataMahasiswa = append(dataMahasiswa, mahasiswa)
		i++
	}

	return dataMahasiswa
}

// TampilkanDataMahasiswa menampilkan data mahasiswa dalam bentuk tabel
func TampilkanDataMahasiswa(dataMahasiswa []Mahasiswa, tampilkanSemua bool) {
	fmt.Println("\n=============================================================================")
	fmt.Println("                            Data Mahasiswa                                   ")
	fmt.Println("=============================================================================")
	fmt.Printf("%-4s %-10s %-20s %-10s %-10s %-10s %-10s\n",
		"NO", "NPM", "Nama", "UTS", "UAS", "Rata-Rata", "Grade")
	fmt.Println("=============================================================================")

	for _, mahasiswa := range dataMahasiswa {
		if tampilkanSemua || mahasiswa.Grade != "X" {
			fmt.Printf("%-4d %-10s %-20s %-10.2f %-10.2f %-10.2f %-10s\n",
				mahasiswa.NO, mahasiswa.NPM, mahasiswa.Nama, mahasiswa.UTS,
				mahasiswa.UAS, mahasiswa.RataRata, mahasiswa.Grade)
		}
	}

	fmt.Println("=============================================================================")
}

// TampilkanStatistik menampilkan statistik dari data mahasiswa
func TampilkanStatistik(dataMahasiswa []Mahasiswa, tampilkanSemua bool) {
	var totalNilai float64
	var nilaiTertinggiUTS, nilaiTertinggiUAS, nilaiTerendahUTS, nilaiTerendahUAS float64
	var jumlahA, jumlahB, jumlahC, jumlahD, jumlahE int

	for _, mahasiswa := range dataMahasiswa {
		if tampilkanSemua || mahasiswa.Grade != "X" {
			totalNilai += mahasiswa.RataRata

			// Memisahkan nilai tertinggi dan terendah untuk UTS dan UAS
			if mahasiswa.UTS > nilaiTertinggiUTS {
				nilaiTertinggiUTS = mahasiswa.UTS
			}
			if mahasiswa.UAS > nilaiTertinggiUAS {
				nilaiTertinggiUAS = mahasiswa.UAS
			}

			if mahasiswa.UTS < nilaiTerendahUTS || nilaiTerendahUTS == 0 {
				nilaiTerendahUTS = mahasiswa.UTS
			}
			if mahasiswa.UAS < nilaiTerendahUAS || nilaiTerendahUAS == 0 {
				nilaiTerendahUAS = mahasiswa.UAS
			}

			switch mahasiswa.Grade {
			case "A":
				jumlahA++
			case "B":
				jumlahB++
			case "C":
				jumlahC++
			case "D":
				jumlahD++
			case "E":
				jumlahE++
			}
		}
	}

	xGradeCount := 0
	for _, mahasiswa := range dataMahasiswa {
		if mahasiswa.Grade == "X" {
			xGradeCount++
		}
	}

	rataRata := totalNilai / float64(len(dataMahasiswa)-xGradeCount)

	fmt.Println("\n=============================================================================")
	fmt.Println("                          Statistik Mahasiswa                                ")
	fmt.Println("=============================================================================")
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata)
	fmt.Printf("Nilai tertinggi UTS: %.2f\n", nilaiTertinggiUTS)
	fmt.Printf("Nilai tertinggi UAS: %.2f\n", nilaiTertinggiUAS)
	fmt.Printf("Nilai terendah UTS: %.2f\n", nilaiTerendahUTS)
	fmt.Printf("Nilai terendah UAS: %.2f\n", nilaiTerendahUAS)
	fmt.Println("=============================================================================")
	fmt.Println("Jumlah mahasiswa dengan grade:")
	fmt.Printf("A: %d\n", jumlahA)
	fmt.Printf("B: %d\n", jumlahB)
	fmt.Printf("C: %d\n", jumlahC)
	fmt.Printf("D: %d\n", jumlahD)
	fmt.Printf("E: %d\n", jumlahE)
	fmt.Println("=============================================================================")
}

// getUserInput membaca input dari pengguna
func getUserInput() string {
	var input string
	fmt.Scan(&input)
	return input
}

// HapusDataMahasiswa menghapus data mahasiswa berdasarkan NPM
func HapusDataMahasiswa(npm string, dataMahasiswa []Mahasiswa) []Mahasiswa {
	for i, mahasiswa := range dataMahasiswa {
		if mahasiswa.NPM == npm {
			// Hapus mahasiswa dari slice
			dataMahasiswa = append(dataMahasiswa[:i], dataMahasiswa[i+1:]...)
			fmt.Printf("Data mahasiswa dengan NPM %s berhasil dihapus.\n", npm)
			return dataMahasiswa
		}
	}
	fmt.Printf("Data mahasiswa dengan NPM %s tidak ditemukan.\n", npm)
	return dataMahasiswa
}

// EditDataMahasiswa mengedit data mahasiswa berdasarkan NPM
func EditDataMahasiswa(npm string, konfig konfigGrade, dataMahasiswa []Mahasiswa) []Mahasiswa {
	for i, mahasiswa := range dataMahasiswa {
		if mahasiswa.NPM == npm {
			// Hapus mahasiswa lama dari slice
			dataMahasiswa = append(dataMahasiswa[:i], dataMahasiswa[i+1:]...)

			// Input data mahasiswa baru
			var mahasiswaBaru Mahasiswa
			mahasiswaBaru.NPM = npm
			fmt.Print("Nama baru: ")
			fmt.Scan(&mahasiswaBaru.Nama)

			// Validasi input nilai UTS
			for {
				fmt.Print("Nilai UTS baru: ")
				nilaiUTS, err := strconv.ParseFloat(getUserInput(), 64)
				if err == nil {
					mahasiswaBaru.UTS = nilaiUTS
					break
				}
				fmt.Println("Error: Input yang Anda masukkan bukan angka.")
			}

			// Validasi input nilai UAS
			for {
				fmt.Print("Nilai UAS baru: ")
				nilaiUAS, err := strconv.ParseFloat(getUserInput(), 64)
				if err == nil {
					mahasiswaBaru.UAS = nilaiUAS
					break
				}
				fmt.Println("Error: Input yang Anda masukkan bukan angka.")
			}

			// Validasi pembagian nol sebelum perhitungan rata-rata
			if mahasiswaBaru.UTS == mahasiswaBaru.UAS {
				mahasiswaBaru.RataRata = mahasiswaBaru.UTS
			} else {
				mahasiswa.RataRata = (mahasiswa.UTS + mahasiswa.UAS) / 2
			}

			mahasiswaBaru.Grade = HitungGrade(mahasiswaBaru.RataRata, konfig)
			mahasiswaBaru.NO = i + 1

			// Tambah mahasiswa baru ke slice
			dataMahasiswa = append(dataMahasiswa, mahasiswaBaru)

			fmt.Printf("Data mahasiswa dengan NPM %s berhasil diedit.\n", npm)
			return dataMahasiswa
		}
	}
	fmt.Printf("Data mahasiswa dengan NPM %s tidak ditemukan.\n", npm)
	return dataMahasiswa
}

// TambahDataMahasiswa menambahkan data mahasiswa baru ke slice
func TambahDataMahasiswa(konfig konfigGrade, dataMahasiswa []Mahasiswa) []Mahasiswa {
	var mahasiswa Mahasiswa

	for {
		fmt.Print("NPM: ")
		fmt.Scan(&mahasiswa.NPM)

		// Validasi NPM unik
		if CekNPMUnik(mahasiswa.NPM, dataMahasiswa) {
			break
		}
		fmt.Println("Error: NPM sudah ada. Harap masukkan NPM yang unik.")
	}

	fmt.Print("Nama: ")
	fmt.Scan(&mahasiswa.Nama)

	// Validasi input nilai UTS
	for {
		fmt.Print("Nilai UTS: ")
		nilaiUTS, err := strconv.ParseFloat(getUserInput(), 64)
		if err == nil {
			mahasiswa.UTS = nilaiUTS
			break
		}
		fmt.Println("Error: Input yang Anda masukkan bukan angka.")
	}

	// Validasi input nilai UAS
	for {
		fmt.Print("Nilai UAS: ")
		nilaiUAS, err := strconv.ParseFloat(getUserInput(), 64)
		if err == nil {
			mahasiswa.UAS = nilaiUAS
			break
		}
		fmt.Println("Error: Input yang Anda masukkan bukan angka.")
	}

	// Validasi pembagian nol sebelum perhitungan rata-rata
	if mahasiswa.UTS == mahasiswa.UAS {
		mahasiswa.RataRata = mahasiswa.UTS
	} else {
		mahasiswa.RataRata = (mahasiswa.UTS + mahasiswa.UAS) / 2
	}

	mahasiswa.Grade = HitungGrade(mahasiswa.RataRata, konfig)
	mahasiswa.NO = len(dataMahasiswa) + 1

	// Tambah mahasiswa ke slice
	dataMahasiswa = append(dataMahasiswa, mahasiswa)

	fmt.Printf("Data mahasiswa dengan NPM %s berhasil ditambahkan.\n", mahasiswa.NPM)

	return dataMahasiswa
}

// UrutkanData mengurutkan data mahasiswa berdasarkan kriteria tertentu
func UrutkanData(dataMahasiswa []Mahasiswa, opsiUrut string) []Mahasiswa {
	switch opsiUrut {
	case "1":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].NPM < dataMahasiswa[j].NPM
		})
	case "2":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].NPM > dataMahasiswa[j].NPM
		})
	case "3":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].Nama < dataMahasiswa[j].Nama
		})
	case "4":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].Nama > dataMahasiswa[j].Nama
		})
	case "5":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].UTS < dataMahasiswa[j].UTS
		})
	case "6":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].UTS > dataMahasiswa[j].UTS
		})
	case "7":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].UAS < dataMahasiswa[j].UAS
		})
	case "8":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].UAS > dataMahasiswa[j].UAS
		})
	case "9":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].RataRata < dataMahasiswa[j].RataRata
		})
	case "10":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].RataRata > dataMahasiswa[j].RataRata
		})
	case "11":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].Grade < dataMahasiswa[j].Grade
		})
	case "12":
		sort.Slice(dataMahasiswa, func(i, j int) bool {
			return dataMahasiswa[i].Grade > dataMahasiswa[j].Grade
		})
	default:
		fmt.Println("Pilihan pengurutan tidak valid.")
	}

	return dataMahasiswa
}

// ExportDataToCSV mengekspor data mahasiswa ke file CSV
func ExportDataToCSV(dataMahasiswa []Mahasiswa, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"NO", "NPM", "Nama", "UTS", "UAS", "Rata-Rata", "Grade"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, mahasiswa := range dataMahasiswa {
		record := []string{
			strconv.Itoa(mahasiswa.NO),
			mahasiswa.NPM,
			mahasiswa.Nama,
			strconv.FormatFloat(mahasiswa.UTS, 'f', 2, 64),
			strconv.FormatFloat(mahasiswa.UAS, 'f', 2, 64),
			strconv.FormatFloat(mahasiswa.RataRata, 'f', 2, 64),
			mahasiswa.Grade,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	fmt.Printf("Data mahasiswa berhasil diekspor ke file %s\n", fileName)
	return nil
}

// CariDataMahasiswa melakukan pencarian data mahasiswa berdasarkan NPM atau Nama
func CariDataMahasiswa(dataMahasiswa []Mahasiswa, kriteria string) []Mahasiswa {
	var hasilPencarian []Mahasiswa

	for _, mahasiswa := range dataMahasiswa {
		if strings.Contains(mahasiswa.NPM, kriteria) || strings.Contains(mahasiswa.Nama, kriteria) {
			hasilPencarian = append(hasilPencarian, mahasiswa)
		}
	}

	return hasilPencarian
}

func main() {
	Header()
	var jmlMahasiswa int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jmlMahasiswa)

	// Validasi jumlah mahasiswa
	if jmlMahasiswa <= 0 {
		fmt.Println("Error: Jumlah mahasiswa harus lebih dari 0.")
		os.Exit(1)
	}

	var konfig konfigGrade
	konfig.BatasA = 80.0
	konfig.BatasB = 70.0
	konfig.BatasC = 60.0
	konfig.BatasD = 50.0
	konfig.BatasMin = 0.0

	var dataMahasiswa []Mahasiswa

	dataMahasiswa = InputDataMahasiswa(jmlMahasiswa, konfig, dataMahasiswa)

	for {
		fmt.Println("\nPilihan Tampilkan Data:")
		fmt.Println("1. Tampilkan Data Keseluruhan (Termasuk Grade X)")
		fmt.Println("2. Tampilkan Data Filter (Tanpa Grade X)")
		fmt.Println("3. Hapus Data Mahasiswa")
		fmt.Println("4. Edit Data Mahasiswa")
		fmt.Println("5. Tambah Data Mahasiswa Baru")
		fmt.Println("6. Urutkan Data")
		fmt.Println("7. Pencarian Data")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih: ")

		pilihan := getUserInput()

		switch pilihan {
		case "1":
			TampilkanDataMahasiswa(dataMahasiswa, true)
			TampilkanStatistik(dataMahasiswa, false)
		case "2":
			TampilkanDataMahasiswa(dataMahasiswa, false)
			TampilkanStatistik(dataMahasiswa, false)
		case "3":
			fmt.Print("Masukkan NPM mahasiswa yang akan dihapus: ")
			npm := getUserInput()
			dataMahasiswa = HapusDataMahasiswa(npm, dataMahasiswa)
		case "4":
			fmt.Print("Masukkan NPM mahasiswa yang akan diedit: ")
			npm := getUserInput()
			dataMahasiswa = EditDataMahasiswa(npm, konfig, dataMahasiswa)
		case "5":
			dataMahasiswa = TambahDataMahasiswa(konfig, dataMahasiswa)
		case "6":
			fmt.Println("\nPilihan Urutkan Data:")
			fmt.Println("1. NPM (Ascending)")
			fmt.Println("2. NPM (Descending)")
			fmt.Println("3. Nama (Ascending)")
			fmt.Println("4. Nama (Descending)")
			fmt.Println("5. UTS (Ascending)")
			fmt.Println("6. UTS (Descending)")
			fmt.Println("7. UAS (Ascending)")
			fmt.Println("8. UAS (Descending)")
			fmt.Println("9. Rata-Rata (Ascending)")
			fmt.Println("10. Rata-Rata (Descending)")
			fmt.Println("11. Grade (Ascending)")
			fmt.Println("12. Grade (Descending)")
			fmt.Print("Pilih opsi pengurutan: ")

			opsiUrut := getUserInput()
			dataMahasiswa = UrutkanData(dataMahasiswa, opsiUrut)
			TampilkanDataMahasiswa(dataMahasiswa, false)

		case "7":
			fmt.Print("Masukkan NPM atau Nama mahasiswa untuk pencarian: ")
			kriteriaPencarian := getUserInput()
			hasilPencarian := CariDataMahasiswa(dataMahasiswa, kriteriaPencarian)
			TampilkanDataMahasiswa(hasilPencarian, false)
			TampilkanStatistik(hasilPencarian, false)
		case "8":
			fmt.Println("Progam Made By Yoga Ardiansyah")
			fmt.Println("Made with Akai Haato ❤️")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih kembali.")
		}
	}
}

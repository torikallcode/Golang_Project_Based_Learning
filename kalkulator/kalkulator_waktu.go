package kalkulator

import "fmt"

func Waktu() {

	var value, value_jam, result float64
	var satuan_awal, satuan_akhir string

	for {
		fmt.Println("=== Selamat datang ===")
		fmt.Print("Masukkan nilai: ")
		fmt.Scanln(&value)
		fmt.Print("Masukkan satuan awal (detik, menit, jam): ")
		fmt.Scanln(&satuan_awal)
		fmt.Print("Masukkan satuan akhir (detik, menit, jam): ")
		fmt.Scanln(&satuan_akhir)

		switch satuan_awal {
		case "jam":
			value_jam = value
		case "menit":
			value_jam = value / 60
		case "detik":
			value_jam = value / 360
		default:
			fmt.Println("Satuan tidak ditemukan")
			return
		}

		switch satuan_akhir {
		case "jam":
			result = value_jam
		case "menit":
			result = value_jam * 60
		case "detik":
			result = value_jam * 360
		default:
			fmt.Println("Satuan tidak ditemukan")
			return
		}

		fmt.Printf("%v %v = %v %v \n", value, satuan_awal, result, satuan_akhir)
	}
}

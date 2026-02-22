package kalkulator

import "fmt"

func Panjang() {

	var satuan_awal, satuan_akhir string
	var value float64

	for {
		fmt.Println("=== Selamat datang di kalkulator massa ===")
		fmt.Print("Masukkan nilai: ")
		fmt.Scanln(&value)
		fmt.Print("Masukkan satuan awal (km, hm, dam, m, dm, cm, mm): ")
		fmt.Scanln(&satuan_awal)
		fmt.Print("Masukkan satuan akhir (km, hm, dam, m, dm, cm, mm): ")
		fmt.Scanln(&satuan_akhir)

		var value_kg float64

		switch satuan_awal {
		case "km":
			value_kg = value
		case "hm":
			value_kg = value / 10
		case "dam":
			value_kg = value / 100
		case "m":
			value_kg = value / 1000
		case "dm":
			value_kg = value / 10000
		case "cm":
			value_kg = value / 100000
		case "mm":
			value_kg = value / 1000000
		default:
			fmt.Println("Satuan tidak ditemukan")
			return
		}

		var result float64

		switch satuan_akhir {
		case "km":
			result = value_kg
		case "hm":
			result = value_kg * 10
		case "dam":
			result = value_kg * 100
		case "m":
			result = value_kg * 1000
		case "dm":
			result = value_kg * 10000
		case "cm":
			result = value_kg * 100000
		case "mm":
			result = value_kg * 1000000
		default:
			fmt.Println("Satuan tidak ditemukan")
			return
		}

		fmt.Printf("%v %v = %v %v \n", value, satuan_awal, result, satuan_akhir)
	}
}

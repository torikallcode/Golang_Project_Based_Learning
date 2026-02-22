package kalkulator

import "fmt"

func Massa() {

	var satuan_awal, satuan_akhir string
	var value float64

	for {
		fmt.Println("=== Selamat datang di kalkulator massa ===")
		fmt.Print("Masukkan nilai: ")
		fmt.Scanln(&value)
		fmt.Print("Masukkan satuan awal (kg, hg, dag, g, dg, cg, mg): ")
		fmt.Scanln(&satuan_awal)
		fmt.Print("Masukkan satuan akhir (kg, hg, dag, g, dg, cg, mg): ")
		fmt.Scanln(&satuan_akhir)

		var value_kg float64

		switch satuan_awal {
		case "kg":
			value_kg = value
		case "hg":
			value_kg = value / 10
		case "dag":
			value_kg = value / 100
		case "g":
			value_kg = value / 1000
		case "dg":
			value_kg = value / 10000
		case "cg":
			value_kg = value / 100000
		case "mg":
			value_kg = value / 1000000
		default:
			fmt.Println("Satuan tidak ditemukan")
			return
		}

		var result float64

		switch satuan_akhir {
		case "kg":
			result = value_kg
		case "hg":
			result = value_kg * 10
		case "dag":
			result = value_kg * 100
		case "g":
			result = value_kg * 1000
		case "dg":
			result = value_kg * 10000
		case "cg":
			result = value_kg * 100000
		case "mg":
			result = value_kg * 1000000
		default:
			fmt.Println("Satuan tidak ditemukan")
			return
		}

		fmt.Printf("%v %v = %v %v \n", value, satuan_awal, result, satuan_akhir)
	}
}

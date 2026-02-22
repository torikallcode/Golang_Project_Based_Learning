package kalkulator

import (
	"bufio"
	"fmt"
	"os"
)

func rumus_diskon(harga, diskon float64) float64 {
	diskon = diskon / 100
	harga_akhir := harga * diskon
	return harga_akhir
}

func lanjut_diskon() bool {

	var pilihan string

	fmt.Print("Lanjut menghitung? (y/n): ")
	fmt.Scanln(&pilihan)
	return pilihan == "y" || pilihan == "Y"
}

func Diskon() {

	reader := bufio.NewReader(os.Stdin)

	var harga, diskon float64

	for {
		fmt.Println("=== Selamat datang di Kalkulator Diskon ===")
		fmt.Print("Masukkan harga awal: ")
		fmt.Fscanln(reader, &harga)
		fmt.Print("Masukkan diskon (%): ")
		fmt.Fscanln(reader, &diskon)
		fmt.Printf("%v diskon %v persen = %v \n", harga, diskon, rumus_diskon(harga, diskon))
		if !lanjut() {
			fmt.Println("Terimakasih telah menggunakan kalkulator diskon")
			break
		}
	}
}

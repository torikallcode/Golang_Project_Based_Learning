package main

import (
	"bufio"
	"calculator/kalkulator"
	"fmt"
	"os"
)

func lanjut() bool {
	var pilihan string
	fmt.Println("lanjut menggunakan kalkulator? (y/n): ")
	fmt.Scanln(&pilihan)
	return pilihan == "y" || pilihan == "Y"
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var pilihan int

	for {
		fmt.Println("=== Selamat Datang ===")
		fmt.Print("1. Kalkulator Sederhana \n2. Kalkulator BMI \n3. Kalkulator Suhu \n4. Kalkulator Diskon \n5.Kalkulator massa \n")
		fmt.Print("Pilih Kalkulator: ")
		fmt.Fscanln(reader, &pilihan)
		switch pilihan {
		case 1:
			kalkulator.Kalkulator_sederhana()
		case 2:
			kalkulator.Bmi()
		case 3:
			kalkulator.Suhu()
		case 4:
			kalkulator.Diskon()
		case 5:
			kalkulator.Massa()
		}
	}

}

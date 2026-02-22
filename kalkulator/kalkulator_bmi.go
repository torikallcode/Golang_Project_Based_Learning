package kalkulator

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func calc(tinggi, berat float64) float64 {
	tinggi = tinggi / 100
	tinggi = math.Pow(tinggi, 2)
	return berat / tinggi
}

func lanjut_bmi() bool {
	var pilihan string
	fmt.Print("Lanjut menghitung BMI? (y/n): ")
	fmt.Scanln(&pilihan)
	return pilihan == "y" || pilihan == "Y"
}

func Bmi() {

	reader := bufio.NewReader(os.Stdin)

	var tinggi, berat float64

	for {
		fmt.Println("=== Selamat datang di kalkulator BMI ===")
		fmt.Print("Masukkan tinggi badan: ")
		fmt.Fscanln(reader, &tinggi)
		fmt.Print("Masukkan berat badan: ")
		fmt.Fscanln(reader, &berat)
		result := calc(tinggi, berat)
		fmt.Printf("Hasil: %.2f\n", result)
		if !lanjut_bmi() {
			fmt.Println("Terimakasih telah menggunakan kalkulator BMI")
			break
		}
	}

}

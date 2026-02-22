package kalkulator

import (
	"bufio"
	"fmt"
	"os"
)

func lanjut_suhu() bool {
	var pilihan string
	fmt.Print("Lanjut menghitung? (y/n): ")
	fmt.Scanln(&pilihan)
	return pilihan == "y" || pilihan == "Y"
}

func Suhu() {

	reader := bufio.NewReader(os.Stdin)
	var pilihan int
	var a float64

	for {
		fmt.Println("=== Selamat Datang di kalkulator Suhu ===")
		fmt.Println("Pilih Satuan suhu awal")
		fmt.Print("1. Celcius \n2. Fahrenheit \n3. Kelvin\n")
		fmt.Print("Masukkan pilihan: ")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:
			var pilihan int
			fmt.Print("Masukkan nilai suhu: ")
			fmt.Fscanln(reader, &a)
			fmt.Print("1. Fahrenheit \n2. Kelvin\n")
			fmt.Print("Konversi satuan ke: ")
			fmt.Scanln(&pilihan)
			switch pilihan {
			case 1:
				result := func(a float64) float64 {
					return (a * 9 / 5) + 32
				}
				fmt.Printf("%.2f Celcius = %2.f Fahrenheit\n", a, result(a))
			case 2:
				result := func(a float64) float64 {
					return a + 275.15
				}
				fmt.Printf("%.2f Celcius = %2.f Kelvin\n", a, result(a))
			}
		case 2:
			var pilihan int
			fmt.Print("Masukkan nilai suhu: ")
			fmt.Fscanln(reader, &a)
			fmt.Print("1. Celcius \n2. Kelvin\n")
			fmt.Print("Konversi satuan ke: ")
			fmt.Scanln(&pilihan)
			switch pilihan {
			case 1:
				result := func(fahrenheit float64) float64 {
					return (fahrenheit - 32) * 5 / 9
				}
				fmt.Printf("%.2f Fahrenheit = %2.f Celcius\n", a, result(a))
			case 2:
				result := func(fahrenheit float64) float64 {
					return (fahrenheit-32)*5/9 + 273.15
				}
				fmt.Printf("%.2f Fahrenheit = %2.f Kelvin\n", a, result(a))
			}
		case 3:
			var pilihan int
			fmt.Print("Masukkan nilai suhu: ")
			fmt.Fscanln(reader, &a)
			fmt.Print("1. Celcius \n2. Kelvin\n")
			fmt.Print("Konversi satuan ke: ")
			fmt.Scanln(&pilihan)
			switch pilihan {
			case 1:
				result := func(kelvin float64) float64 {
					return kelvin - 273.15
				}
				fmt.Printf("%.2f Kelvin = %2.f Celcius\n", a, result(a))
			case 2:
				result := func(kelvin float64) float64 {
					return (kelvin-273.15)*9/5 + 32
				}
				fmt.Printf("%.2f Kelvin = %2.f Fahrenheit\n", a, result(a))
			}
		}

		if !lanjut_suhu() {
			fmt.Println("Terimakasih...")
			break
		}
	}
}

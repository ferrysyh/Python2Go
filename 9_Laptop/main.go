package main

import "fmt"

func main() {
	var baterai int
	fmt.Print("Masukkan nilai baterai: ")
	fmt.Scan(&baterai)

	if baterai == 100 {
		fmt.Println("Baterai Full")
	} else if baterai >= 75 && baterai <= 99 {
		fmt.Println("Charge dalam 30 menit")
	} else if baterai >= 20 && baterai <= 74 {
		fmt.Println("Charge dalam 10 menit")
	} else {
		fmt.Println("Segera charge")
	}
}

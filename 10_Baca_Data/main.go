package main

import "fmt"

func main() {
	var teks string
	fmt.Print("Masukkan teks: ")
	fmt.Scan(&teks)

	for teks != "SELESAI" {
		fmt.Println(teks)
		fmt.Print("Masukkan teks: ")
		fmt.Scan(&teks)
	}
}

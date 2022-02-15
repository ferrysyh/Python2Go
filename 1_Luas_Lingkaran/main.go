package main

import "fmt"

func main() {
	var r float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)
	
	luas := 3.14 * r * r
	fmt.Println("Luas lingkaran: ", luas)
}

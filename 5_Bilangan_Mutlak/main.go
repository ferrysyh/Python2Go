package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	var hasil int = a - b
	if hasil < 0 {
		hasil = hasil * -1
	}
	fmt.Println("Nilai mutlak dari", a, "dikurangi", b, "adalah", hasil)
}

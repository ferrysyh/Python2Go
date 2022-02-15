package main

import "fmt"

func main() {
	var x1, x2, x3 int
	fmt.Print("Masukkan nilai x1: ")
	fmt.Scan(&x1)
	fmt.Print("Masukkan nilai x2: ")
	fmt.Scan(&x2)
	fmt.Print("Masukkan nilai x3: ")
	fmt.Scan(&x3)

	if x1+x2 > x3 {
		if x2+x3 > x1 {
			if x3+x1 > x2 {
				fmt.Println("Sisi segitiga")
			} else {
				fmt.Println("Bukan sisi segitiga")
			}
		} else {
			fmt.Println("Bukan sisi segitiga")
		}
	} else {
		fmt.Println("Bukan sisi segitiga")
	}
}

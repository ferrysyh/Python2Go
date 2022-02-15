package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	if a < b && b < c {
		fmt.Println("Tiga bilangan terurut")
	} else {
		fmt.Println("Tidak terurut")
	}
}

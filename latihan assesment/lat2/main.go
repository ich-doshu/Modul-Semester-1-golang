package main 

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	modulus := a % b
	fmt.Println(tambah, kurang, kali, bagi, modulus)
}
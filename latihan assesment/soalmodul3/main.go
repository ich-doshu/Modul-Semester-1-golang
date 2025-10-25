package main 

import "fmt"

func main() {
	var n int
	fmt.Println("Masukkan nilai dalam qirat")
	fmt.Scan(&n)
	dinar := n / 600
	dirham := n % 600 / 60
	fals := n %600 %60 /60
	qirat := n % 600 % 60 % 6
	fmt.Println(dinar, dirham, fals, qirat)
}
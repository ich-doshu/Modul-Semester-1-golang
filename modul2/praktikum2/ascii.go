package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5  byte
	var b1, b2, b3 int 
	fmt.Print("masukkan angka : ")
	fmt.Scan(&a1, &a2, &a3, &a4, &a5 )
	fmt.Print("masukkan karakter : ")
	fmt.Scanf("%c", &b1)
	fmt.Scanf("%c", &b2)
	fmt.Scanf("%c", &b3)
	fmt.Printf("%c%c%c%c%c", a1, a2, a3, a4, a5)
	fmt.Printf("%c%c%c", b1+1, b2+1, b3+1)
}
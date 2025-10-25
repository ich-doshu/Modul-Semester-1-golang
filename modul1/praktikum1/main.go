package main

import "fmt"

func main() {
	fmt.Println("Selamat Datang di Telyu")
	nama := "" 
	fmt.Print("Masukkan nami sampean:")
	fmt.Scan(&nama)
	fmt.Printf("Halo, %s", nama)
}
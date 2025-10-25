package main 

import "fmt"

func main() {
	var nama string
	var nim int
	fmt.Println("=====Selamat Datang Member=====")
	fmt.Scan(&nama)
	fmt.Scan(&nim)
	fmt.Println("Nama: ", nama)
	fmt.Print("NIM: ", nim)
}
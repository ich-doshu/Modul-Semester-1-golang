package main

import "fmt"

func main() {
	var x, fx float64
	fmt.Printf("Masukkan angka : ",  )
	fmt.Scanln(&x)
	fx = 2 / (x+5) + 5
	fmt.Println("hasilnya adalah", fx)
}
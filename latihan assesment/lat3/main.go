package main

import "fmt"

func main() {
	var angka1, angka2 int
	fmt.Scan(&angka1)
	fmt.Scan(&angka2)
	for i := angka2; i >= angka1; i--{
		fmt.Print("", i)
	}
}
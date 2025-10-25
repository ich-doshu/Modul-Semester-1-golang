package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	jumlah := 0
	for i := x; i <=y; i+=1 {
		jumlah += i
	}
	fmt.Print(jumlah)
}
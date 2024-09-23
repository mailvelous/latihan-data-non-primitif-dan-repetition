package main

import (
	"fmt"
)

func main() {
	var jumlahUang int
	fmt.Println("Masukkan Jumlah Uang : ")
	fmt.Scan(&jumlahUang)

	var sisaJumlahUangBagi1000 int
	var jumlahPecahanUangBagi1000 int

	jumlahPecahanUangBagi1000 = jumlahUang / 1000
	sisaJumlahUangBagi1000 = jumlahUang % 1000
	fmt.Println("Jumlah Pecahan Uang 1000 : ", jumlahPecahanUangBagi1000)
	fmt.Println("Jumlah Sisa Uang : ", sisaJumlahUangBagi1000)

	// var sisaJumlahUangBagi500 int
	// var jumlahPecahanUang500 int

}

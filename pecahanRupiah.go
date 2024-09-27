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

	var sisaJumlahUangBagi500 int
	var pecahanUangBagi500 int

	pecahanUangBagi500 = sisaJumlahUangBagi1000 / 500
	sisaJumlahUangBagi500 = sisaJumlahUangBagi1000 % 500
	fmt.Println("Jumlah Pecahan Uang 500 : ", pecahanUangBagi500)
	fmt.Println("Jumlah Sisa Uang : ", sisaJumlahUangBagi500)

	var sisaJumlahUangBagi100 int
	var pecahanUangBagi100 int

	pecahanUangBagi100 = sisaJumlahUangBagi500 / 100
	sisaJumlahUangBagi100 = sisaJumlahUangBagi500 % 100
	fmt.Println("Jumlah Pecahan Uang 100 : ", pecahanUangBagi100)
	fmt.Println("Jumlah Sisa Uang : ", sisaJumlahUangBagi100)

	var sisaJumlahUangBagi50 int
	var pecahanUangBagi50 int
	
	pecahanUangBagi50 = sisaJumlahUangBagi100 / 50
	sisaJumlahUangBagi50 = sisaJumlahUangBagi100 % 50
	fmt.Println("Jumlah Pecahan Uang 50 : ", pecahanUangBagi50)
	fmt.Println("Jumlah Sisa Uang : ", sisaJumlahUangBagi50)

	var sisaJumlahUangBagi25 int
	var pecahanUangBagi25 int

	pecahanUangBagi25 = sisaJumlahUangBagi50 /25
	sisaJumlahUangBagi25 = sisaJumlahUangBagi50 % 25
	fmt.Println("Jumlah Pecahan Uang 25 : ", pecahanUangBagi25)
	fmt.Println("Jumlah Sisa Uang : ", sisaJumlahUangBagi25)

}

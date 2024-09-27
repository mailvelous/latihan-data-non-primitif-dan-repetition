package main

import (
	"fmt"
)

func main() {
	hitungAngkaGanjil()

}

func hitungAngkaGanjil() {
	var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// jumlahBilanganGanjil := 0
	jumlahBilanganGanjil := 0

	for _, number := range numbers {
		if number % 2 == 1 {
			fmt.Printf("%d", number)

			fmt.Println()

			jumlahBilanganGanjil++
			fmt.Printf("%d", jumlahBilanganGanjil)
			

		}

		// for i := 0; i <= number; i++{
		// 	fmt.Println(i)
		// }
	}
}
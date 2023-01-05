package main

import "fmt"

const maxArray int = 100

type tabInt [maxArray]int

func pemilih(mArr tabInt, n int)

func main() {
	var calon tabInt
	var masuk, sah, pilihan int
	masuk = 0
	sah = 0
	fmt.Scan(&pilihan)
	for pilihan != 0 {
		masuk++
		if pilihan >= 1 && pilihan <= 20 {
			sah++
			calon[pilihan-1]++
		}
		fmt.Scan(&pilihan)
	}

	fmt.Println("Suara Masuk :", masuk)
	fmt.Println("Suara sah :", sah)

	for i := 0; i < 20; i++ {
		if calon[i] != 0 {
			fmt.Print(i+1, ":", calon[i], "\n")
		}
	}
}

package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	/* buatlah kode utama yang membaca baris pertama (
	   n dan k). kemudian data diisi
	   oleh prosedur isiArray(n), dan pencarian oleh fungsi posisi(n,k), dan setelah
	   itu output dicetak */

	var n, k, hasil int
	fmt.Scan(&n, &k)
	isiArray(n)
	hasil = posisi(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")

	} else {
		fmt.Println(hasil)
	}
}

func isiArray(n int) {
	/* I
	S. terdefinisi integer n, dan sejumlah n d ata sudah siap pada piranti masukan.
	F S. Array data berisi n (<=NMAX) bilangan */
	for i := 0; i < n && i < NMAX; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	/* mengembalikan posisi k dalam array data dengan n elemen. Posisi dimulai dari
	posisi 0. Jika tidak ada kembalikan 1 */
	var i int

	for i < n && data[i] != k {
		i++
	}
	if i < n {
		return i
	} else {
		return -1
	}
}

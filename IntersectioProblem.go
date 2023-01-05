package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	/* mengembalikan true apabila bilangan val ada di dalam
	   array T yang berisi sejumlah n bilangan bulat */
	var i int = 0
	var stop bool
	stop = false

	for i < n && !stop {
		if T[i] == val {
			stop = true
		} else {
			stop = false
		}
		i++
	}
	return stop
}

func inputSet(T *set, n *int) {
	/* I.S. data himpunan telah siap pada piranti masukan
	F.S. array T berisi sejumlah n bilangan bulat yang berasal
	dari masukan (masukan berakhir apabila bilangan ada yang
	duplikat, atau array penuh) Catatan: Panggil function exist
	di sini untuk membantu pengecekan */

	var Input int
	var stop bool
	*n = 0

	fmt.Scan(&Input)
	stop = exist(*T, *n, Input)
	for *n < 2002 && !stop {
		T[*n] = Input
		*n++
		fmt.Scan(&Input)
		stop = exist(*T, *n, Input)
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	/* I.S. terdefinisi himpunan T1 dan T2 yang berisi sejumlah n
	dan m anggota himpunan
	F.S. himpunan T3 berisi sejumlah h bilangan bulat yang
	merupakan irisan dari himpunan T1 dan T2
	Catatan: Panggil function exist di sini untuk
	membantu pengecekan */
	var i int
	for i < m {
		if exist(T1, n, T2[i]) {
			T3[*h] = T2[i]
			*h++
		}
		i++
	}
}

func printSet(T set, n int) {
	/* I.S. terdefinisi sebuah himpunan T yang berisi
	sejumlah n bilangan bulat F.S. menampilkan isi array T
	secara horizontal (dipisahkan oleh spasi) */
	var i int = 0
	var j int = 0
	var k int
	var temp int
	for j < n {
		for k = j; k < n; k++ {
			if T[j] > T[k] {
				temp = T[j]
				T[j] = T[k]
				T[k] = temp
			}
		}
		j++
	}

	for i < n {
		fmt.Print(T[i], " ")
		i++
	}
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int
	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}

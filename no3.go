package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	var x1 int

	fmt.Scan(&x1)
	fmt.Scan(&nData)

	bacaData(&data, nData)
	cetakData(data, nData)

	fmt.Print("Hasil pencarian: ")
	if sequentialSearch(data, nData, x1) {
		jumlah := frekuensiBilangan(data, nData, x1)
		fmt.Printf("Bilangan ditemukan. Terdapat %d bilangan %d.\n", jumlah, x1)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}

}

func bacaData(A *tabInt, n int) {
	for i := 0; i < n && i < NMAX; i++ {
		fmt.Scan(&(*A)[i])
	}
}

func cetakData(A tabInt, n int) {
	fmt.Print("Data Bilangan : ")
	for i := 0; i < n && i < NMAX; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func frekuensiBilangan(A tabInt, n, x int) int {
	frekuensi := 0
	for i := 0; i < n && i < NMAX; i++ {
		if A[i] == x {
			frekuensi++
		}
	}
	return frekuensi
}

func sequentialSearch(A tabInt, n, x int) bool {
	for i := 0; i < n && i < NMAX; i++ {
		if A[i] == x {
			return true
		}
	}
	return false
}

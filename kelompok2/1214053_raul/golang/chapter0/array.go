package main

import (
	"fmt"
)

func main() {
	var names [3]string
	//ini membuat string 3 variable
	names[0] = "RAUL"
	names[1] = "MAHYA"
	names[2] = "KOMARAN"
	//ini isi arraynya
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	//ini isi memprint dari si string yang mau kita running
	//ini khusus untuk string
	var values = [3]int{
		90,
		95,
		80,
	}
	//membuat int 3 variable
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	//perintah untuk mengrun menggunakan println
	fmt.Println(len(names))
	fmt.Println(len(values))
	//fungsi len (Panjang String di Golang Go – Panjang String Untuk mendapatkan panjang String di pemrograman Go, ubah string menjadi array rune, dan teruskan array ini ke fungsi len(). Sebuah string berisi karakter sebagai titik unicode, bukan byte. len (string) mengembalikan jumlah byte dalam string ini, tetapi bukan jumlah karakter.)
	var mahya [1]string
	fmt.Println(len(mahya))
}

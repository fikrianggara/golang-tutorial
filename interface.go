package main

import (
	"fmt"
	"math"
)

// interface merupakan tipe data
// kumpulan definisi method dan tipe data returnya
// bisa digunakan jika sudah ada isinya, yaitu
// object yang memiliki method dengan nama yang sama dengan interface

type hitung interface {
	getLuas() float32
	getKeliling() float32
}

type kotak struct {
	panjang uint
	lebar   uint
}

func (p kotak) getLuas() float32 {
	return float32(p.panjang * p.lebar)
}

func (p kotak) getKeliling() float32 {
	return float32(2*p.panjang + 2*p.lebar)
}

func (p kotak) Ispersegi() bool {
	return p.panjang == p.lebar
}

type lingkaran struct {
	diamater uint
}

func (l lingkaran) getJariJari() float32 {
	return float32(l.diamater) / 2
}
func (l lingkaran) getLuas() float32 {
	return float32(math.Pi * (l.getJariJari() * l.getJariJari()))
}

func (l lingkaran) getKeliling() float32 {
	return 2 * math.Pi * l.getJariJari()
}

func initInterfaceExample() {
	var persegiObj hitung
	persegiObj = kotak{panjang: 12, lebar: 10}
	var lingkaranObj hitung
	lingkaranObj = lingkaran{10}

	// karena bisa saja ada banyak variabel yang menggunakan tipe data interface,
	// ketika ada method yang dimiliki objek tertentu, kita perlu me-casting sesuai objek yang memiliki method tersebut
	// contoh di bawah adalah cara bagaimana kita mengakses method Ispersegi().

	fmt.Println("apakah persegi :", persegiObj.(kotak).Ispersegi())
	fmt.Println("jari jari lingkaran", lingkaranObj.(lingkaran).getJariJari())
}

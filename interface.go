package main

import (
	"fmt"
	"math"
)

// interface merupakan tipe data
// kumpulan definisi method dan tipe data returnya
// bisa digunakan jika sudah ada isinya, yaitu
// object yang memiliki method dengan nama yang sama dengan interface

type hitung2D interface {
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
	var persegiObj hitung2D
	persegiObj = kotak{panjang: 12, lebar: 10}
	var lingkaranObj hitung2D
	lingkaranObj = lingkaran{10}

	// karena bisa saja ada banyak variabel yang menggunakan tipe data interface,
	// ketika ada method yang dimiliki objek tertentu, kita perlu me-casting sesuai objek yang memiliki method tersebut
	// contoh di bawah adalah cara bagaimana kita mengakses method Ispersegi().

	fmt.Println("apakah persegi :", persegiObj.(kotak).Ispersegi())
	fmt.Println("jari jari lingkaran", lingkaranObj.(lingkaran).getJariJari())
}

// contoh embed interface (interface di dalam interace)
type hitung3D interface {
	volume() float32
}

//kubus inherit dari kotak
type kubus struct {
	kotak
	tinggi uint
}

// inteface di dalam interface
// jika salah satu method di interface butuh argumen pointer, maka inisiasi objek harus mengembalikan alamat
// contoh
// var kubusX hitung = &kubus{}

type hitung interface {
	hitung2D
	hitung3D
}

func (k kubus) volume() float32 {
	return float32(k.panjang * k.lebar * k.tinggi)
}

//tapi kalo instance memiliki method yang tidak ada di definisi interface (contohnya setTinggi())
// tidak perlu menginisiasi objek sebagai alammat

func (k kubus) setTinggi(newT uint) {
	k.tinggi = newT
}

// func (k *kubus)

func initEmbedInterfaceExample() {
	// kotak1 := kotak{panjang: 5, lebar: 4}
	var kubusInterface hitung = kubus{kotak: kotak{panjang: 5, lebar: 4}, tinggi: 6}
	var kubus1 kubus = kubusInterface.(kubus)

	fmt.Println("lebar :", kubus1.lebar)
	fmt.Println("panjang :", kubus1.panjang)
	fmt.Println("tinggi :", kubus1.tinggi)
	fmt.Println("volume:", kubus1.volume())
	kubus1.setTinggi(3)
	fmt.Println("tinggi ", kubus1.tinggi)
	fmt.Println("volume:", kubus1.volume())
}

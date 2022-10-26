package main

import (
	"fmt"
	"os"
	"strconv"
)

type Batch3 struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	// name := "Rifki"
	// age := 20

	// argsRaw := os.Args

	/*
		fmt.Println("Nama Saya Adalah ==>", name)
		fmt.Println("Umur Saya ==>", age)
		fmt.Println("Hello world")
	*/

	// fmt.Printf("%T, %T", name, age)
	// fmt.Printf("-> %#v\n", argsRaw)

	// args := argsRaw[1:]
	// fmt.Printf("-> %#v\n", args)

	var peserta = []Batch3{
		{nama: "RIfki", alamat: "Bogor", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Udin", alamat: "Bali", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Sapri", alamat: "Bandung", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Iyan", alamat: "Jakarta", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Ragil", alamat: "Jogjakarta", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Nabillah", alamat: "Surabaya", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Randi", alamat: "BSD", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Doni", alamat: "Lampung", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Budi", alamat: "Aceh", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
		{nama: "Ando", alamat: "Pamulang", pekerjaan: "Mahasiswa", alasan: "Menambahilmu"},
	}

	arg, err := strconv.Atoi(os.Args[1:][0])
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("Nama \t\t= ", peserta[arg-1].nama)
	fmt.Println("Alamat \t\t= ", peserta[arg-1].alamat)
	fmt.Println("pekerjaan \t= ", peserta[arg-1].pekerjaan)
	fmt.Println("Alasan \t\t= ", peserta[arg-1].alasan)

}

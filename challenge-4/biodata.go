package main

import (
	"fmt"
	"os"
	"strconv"
)

type Biodata struct {
	nama, alamat, pekerjaan, alasan string
}

func main() {
	params := os.Args[1]
	var data map[Biodata]int = setData()
	angka, _ := strconv.Atoi(params)
	for str, i := range data {
		if i == angka {
			fmt.Println("Nama :", str.nama)
			fmt.Println("Alamat :", str.alamat)
			fmt.Println("Pekerjaan :", str.pekerjaan)
			fmt.Println("Alasan memilih kelas Golang :", str.alasan)
		}
	}

}

func setData() map[Biodata]int {
	a1 := Biodata{nama: "adnan", alamat: "sidoarjo", pekerjaan: "wiraswasta", alasan: "belajar"}
	a2 := Biodata{nama: "pebry", alamat: "cibinong", pekerjaan: "programmer", alasan: "bekerja"}
	a3 := Biodata{nama: "adhit", alamat: "indramayu", pekerjaan: "data analyst", alasan: "memahami dunia"}
	a4 := Biodata{nama: "yudha", alamat: "tangerang", pekerjaan: "data science", alasan: "meneladani sikap"}
	return map[Biodata]int{a1: 1, a2: 2, a3: 3, a4: 4}
}

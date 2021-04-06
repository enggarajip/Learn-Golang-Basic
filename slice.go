package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var bulan = months[4:7]
	fmt.Println(bulan)
	fmt.Println(len(bulan))
	fmt.Println(cap(bulan))

	//months[5] = "diubah"
	//fmt.Println(bulan)

	//bulan[0] = "Mei lagi"
	//fmt.Println(months)

	var bulan1 = months[10:]
	fmt.Println(bulan1)

	var bulan2 = append(bulan1, "Enggar")
	fmt.Println(bulan2)
	bulan2[1] = "Bukan Desember"
	fmt.Println(bulan2)

	fmt.Println(bulan1)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Enggar"
	newSlice[1] = "Aji"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//iniArray := [5]int{1, 2, 3, 4, 5}
	//iniSlice := []int{1, 2, 3, 4, 5}
}

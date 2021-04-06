package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "enggar"
	names[1] = "aji"
	names[2] = "prasetyo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		97,
		98,
		99,
	}
	fmt.Println(values[0])

	var nilaiUjian [10]int // mengisi nilai pada array
	nilaiUjian[5] = 6
	nilaiUjian[8] = 9
	fmt.Println(nilaiUjian[5])
	fmt.Println(nilaiUjian[8])

	fmt.Println(len(nilaiUjian))
	fmt.Println(len(names))

}

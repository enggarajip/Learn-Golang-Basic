package main

import "fmt"

func main() {
	var name1 = "ajip"
	var name2 = "ajip"

	var tes bool = name1 == name2
	fmt.Println(tes)

	var value1 = 1000
	var value2 = 2000

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}

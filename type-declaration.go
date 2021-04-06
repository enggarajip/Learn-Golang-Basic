package main

import "fmt"

func main() {
	type NoKTP string
	type kenyang bool

	var noKTPajip NoKTP = "12345666"
	fmt.Println(noKTPajip)

	var sudahMakan kenyang = true
	fmt.Println(sudahMakan)
}

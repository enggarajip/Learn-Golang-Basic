package main

import "fmt"

func main() {
	var name string

	name = "Enggar Aji"
	fmt.Println(name)

	name = "Enggar AjiP" //mengganti isi variable
	fmt.Println(name)

	var friendName = "Bemo"
	fmt.Println(friendName)

	var age = 26
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	country = "Japan"
	fmt.Println(country)

	var (
		firstName = "Aji"
		lastName  = "Prasetyo"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}

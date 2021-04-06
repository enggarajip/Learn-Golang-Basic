package main

import "fmt"

func main() {
	const firstname string = "Enggar"
	const lastname = "Aji"
	const value = 1000

	fmt.Println(firstname)
	fmt.Println(lastname)
	fmt.Println(value)

	const (
		firstName = "Farah"
		lastName  = "Aulia"
		age       = 13
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
}

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}

	for counter := 1; counter <= 15; counter++ {
		fmt.Println("Perulangan ke", counter)
	} // init statment, post statment

	slice := []string{"enggar", "aji", "prasetyo", "budi", "bagus"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Aji"
	person["title"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

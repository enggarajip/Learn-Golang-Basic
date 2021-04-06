package main

import "fmt"

func main() {
	name := "aji"

	if name == "aji" {
		fmt.Println("Hello ajip")
	} else if name == "enggar" {
		fmt.Println("hello enggar")
	} else if name == "bagas" {
		fmt.Println("hello bagas")
	} else {
		fmt.Println("hi, boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("terlalu panjang")
	}
}

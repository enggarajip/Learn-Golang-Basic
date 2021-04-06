package main

import "fmt"

func main() {
	name := "ajip"

	switch name {
	case "ajip":
		fmt.Println("Halo, ajip!")
	case "enggar":
		fmt.Print("halo, enggar!")
	default:
		fmt.Println("halo, boleh kenalan?")

	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah pas")
	}
}

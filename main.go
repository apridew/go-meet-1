package main

import (
	"fmt"
)

func Panggil() {
	fmt.Println("Halo, salam kenal !")
}

func LuasSegitiga(alas, tinggi float32) float32 {
	hasil := (alas * tinggi) / 2
	return hasil
}

func main() {
	var firstName string = "Dewi"
	var lastName string = "Apri"
	lastName = "Putra"

	nilai := 88

	fmt.Println("Hello World")
	fmt.Println(firstName, lastName)
	fmt.Println("Nilai :", nilai)

	Panggil()
	Panggil()
	Panggil()

	fmt.Println(LuasSegitiga(11, 3))
}

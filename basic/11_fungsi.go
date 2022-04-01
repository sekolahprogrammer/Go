package main

import "fmt"

// tanpa parameter
func hallo() {
	fmt.Println("Hallo Golang!")
}

// dengan parameter
func data(nama string, umur int) {
	fmt.Println("nama : ", nama, ", umur : ", umur)
}

func main() {
	// tanpa parameter
	hallo()

	// dengan parameter
	data("prayogaea", 21)
	data("putriajeng", 27)
}

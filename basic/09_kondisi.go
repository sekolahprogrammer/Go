package main

import "fmt"

func main() {
	var cuaca string = "dingin"

	// kondisi if
	if cuaca == "hujan" {
		fmt.Println("pakai payung")
	} else if cuaca == "dingin" {
		fmt.Println("pakai jaket")
	} else {
		fmt.Println("tidak pakai payung dan jaket")
	}

	//switch
	var hari = 5

	switch hari {
	case 1:
		fmt.Println("senin")
	case 2:
		fmt.Println("selasa")
	case 3:
		fmt.Println("rabu")
	case 4:
		fmt.Println("kamis")
	case 5:
		fmt.Println("jum'at")
	case 6:
		fmt.Println("saptu")
	case 7:
		fmt.Println("minggu")
	default:
		fmt.Println("hari tidak ditemukan")
	}
}

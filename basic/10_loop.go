package main

import "fmt"

func main() {
	// for
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// continue
	for j := 0; j <= 10; j++ {
		if j == 5 {
			continue
		}
		fmt.Println(j)
	}

	//break
	for q := 0; q <= 10; q++ {
		if q == 5 {
			break
		}
		fmt.Println(q)
	}
}

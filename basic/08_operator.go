package main

import "fmt"

var x = 10
var y = 3

func main() {
	// aritmatika
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)

	// penugasan
	x += 5
	fmt.Println(x)
	x -= 5
	fmt.Println(x)
	x *= 5
	fmt.Println(x)
	x /= 5
	fmt.Println(x)
	x %= 3
	fmt.Println(x)

	// perbandingan
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x > y)
	fmt.Println(x < y)
	fmt.Println(x >= y)
	fmt.Println(x <= y)

	//logika
	fmt.Println(x > 5 && x < 20)
	fmt.Println(x > y || y > x)
	fmt.Println(!(x > 5 && x < 20))
}

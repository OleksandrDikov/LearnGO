package main

import "fmt"

const value = 20

func main() {
	var i int = value
	var f float64 = value
	fmt.Println(i, f)

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	fmt.Println(b, smallI, bigI)

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}

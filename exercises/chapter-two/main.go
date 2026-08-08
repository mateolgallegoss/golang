package main

import "fmt"

func main() {

	// Exercise 1
	i := 20
	var f float32 = float32(i)

	fmt.Println("------------------------ EXERCISE 1 ------------------------")
	fmt.Printf("Integer value: %d \n", i)
	fmt.Printf("Float value: %f \n", f)

	// Exercise 2
	const value = 32
	var i2 int = value
	var f2 float64 = value

	fmt.Println("------------------------ EXERCISE 2 ------------------------")
	fmt.Printf("Integer value: %d \n", i2)
	fmt.Printf("Float value: %f \n", f2)

	// Exercise 3
	var b byte
	var smallI int32
	var bigI uint64

	b, smallI, bigI = 128, 65536, 9223372036854775808
	b += 1
	smallI += 1
	bigI += 1

	fmt.Println("------------------------ EXERCISE 3 ------------------------")
	fmt.Printf("Byte value: %d \n", b)
	fmt.Printf("Small Integer value: %d \n", smallI)
	fmt.Printf("Big Integer value: %d \n", bigI)
}

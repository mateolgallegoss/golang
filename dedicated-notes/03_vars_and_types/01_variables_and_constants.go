package main

import "fmt"

// Package-level scope constant
const AppName = "GoLearning"

// iota example
const (
	// iota starts at 0 and increments per line
	StatusPending = iota // 0
	StatusApproved       // 1
	StatusRejected       // 2
	StatusCancelled      // 3
)

const (
	_  = iota             // ignore 0
	KB = 1 << (10 * iota) // 1 << (10*1) = 1024
	MB                    // 1 << (10*2)
	GB                    // 1 << (10*3)
)

func main() {
	// =============================================
	// 1. var vs := (short declaration)
	// =============================================

	// var declares a variable with an explicit type
	var name string = "Alice"
	var age int
	age = 30

	// := is shorthand — type inference, only inside functions
	city := "New York"
	country := "USA"

	fmt.Println("=== var vs :=")
	fmt.Println(name, age, city, country)

	// Multiple variables
	var x, y int = 1, 2
	a, b := "hello", true

	fmt.Println(x, y, a, b)

	// =============================================
	// 2. Zero Values
	// =============================================
	// Every variable without an explicit initial value gets a zero value

	var zeroInt int       // 0
	var zeroFloat float64 // 0.0
	var zeroBool bool     // false
	var zeroStr string    // "" (empty)
	var zeroPtr *int      // nil

	fmt.Println("\n=== Zero Values")
	fmt.Printf("int: %d\nfloat64: %f\nbool: %t\nstring: %q\npointer: %v\n",
		zeroInt, zeroFloat, zeroBool, zeroStr, zeroPtr)

	// =============================================
	// 3. const and iota
	// =============================================

	const Pi = 3.14159
	const Greeting = "Hello, Go!"

	fmt.Println("\n=== Constants")
	fmt.Println(Pi, Greeting)

	fmt.Println("Status codes:", StatusPending, StatusApproved, StatusRejected, StatusCancelled)
	fmt.Println("Sizes: KB=", KB, "MB=", MB, "GB=", GB)

	// =============================================
	// 4. Scope and Shadowing
	// =============================================

	// This shadows the package-level AppName
	AppName := "LocalScope"
	fmt.Println("\n=== Scope & Shadowing")
	fmt.Println("Package-level AppName is accessible only via the package")
	fmt.Println("Local AppName shadows it:", AppName)

	// Block scope
	outer := "outer"
	{
		inner := "inner"
		outer := "shadowed outer" // shadows outer variable above
		fmt.Println("Inside block:", outer, inner)
	}
	fmt.Println("Outside block:", outer) // original outer is unchanged
	// fmt.Println(inner) // error: undefined — inner is out of scope
}

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Literales en go")
	fmt.Println()

	// Literales enteros
	fmt.Print(0B_1000011, " ")
	fmt.Print(0O_505, " ")
	fmt.Print(0X_3f7a, " ")
	fmt.Println(49)
	fmt.Println()

	// Literales de punto flotante
	fmt.Print(12.234e7, " ")
	fmt.Print(158.97E-33, " ")
	fmt.Print(5382e44, " ")
	fmt.Print(0X_e3afP10, "")
	fmt.Print(0X_3d44ff.b3P-12, " ")
	fmt.Println(0X_5ee7.ffa3P12)
	fmt.Println()

	// Literales imaginarios
	fmt.Print(3.1415e10i, " ")
	fmt.Print(158.93e-20i, " ")
	fmt.Print(532e7i, " ")
	fmt.Print(0X_4f3aP4i, " ")
	fmt.Print(0X_3d3a.54P10i, " ")
	fmt.Println(0X_5ee7P12i)
	fmt.Println()

	// Literales rúnicos


	// Literales string
	fmt.Print(`quierokeke`, " ")
	fmt.Print(`ouisdufhdo`, " ")
	fmt.Print(`\abc`, " ")
	fmt.Print(`\geakx`, " ")
}
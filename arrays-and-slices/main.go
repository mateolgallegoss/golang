package main

import (
	"fmt"
	"slices"
)

func main() {
	ar1 := [5]int{4, 3, 1, 0, 5}
	var ar2 = [13]int{7, 6, 40, 3, 4: 1, 3, 10}

	fmt.Println("------------------------ ARRAYS (underrated bullshit) ------------------------")
	fmt.Printf("Primer arreglo: %d \n", ar1)
	fmt.Printf("Segundo arreglo: %d \n", ar2)
	
	slc1 := []int{4, 2, 0, 1}
	var slc2 = []int{}
	slc2 = append(slc2, 3, 4, 7)
	var slc3 = []int{4, 2, 0, 1}
	var slc4 = []string{"glazewm", "btop", "yasb", "windhawk"}
	var slc5 = []float64{4e34, 2.2e6, 3945.1e2, 40.30291}
	var slc6 = []float64{4e30, 12.3e6, 2.2323, 25.67e60}
	

	fmt.Println("------------------------ SLICES (The side goat) ------------------------")

	fmt.Println("SLICES INDEXES")
	fmt.Printf("Primer slice: %d \n", slc1)
	fmt.Printf("Segundo slice: %d \n", slc2)
	fmt.Printf("Tercer slice: %d \n", slc3)
	fmt.Printf("Cuarto slice: %s \n", slc4)
	fmt.Printf("Quinto slice: %f \n", slc5)
	fmt.Printf("Sexto slice: %f \n", slc6)

	fmt.Println("COMPARACIÓN DE IGUALDAD ENTRE SLICES")
	fmt.Printf("El primer slice es igual al segundo: %b \n", slices.Equal(slc1, slc2))
	fmt.Printf("El primer slice es igual al tercero: %b \n", slices.Equal(slc1, slc3))
	fmt.Printf("El segundo slice es igual al tercero: %b \n", slices.Equal(slc2, slc3))
	fmt.Printf("El quinto slice es igual al sexto: %b \n", slices.Equal(slc5, slc6))
}

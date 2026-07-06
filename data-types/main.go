package main

import // Cualquier paquete importado que no se use, devuelve un error
(
	"fmt"
	"strings"
)

// Para variables públicas usar PascalCase
var EnteroUno int = 10

func main() {
	// Cualquier variable declarada que no se use, devuelve un error

	// Para variables privadas usar camelCase

	// NÚMEROS:
	entero := 10
	decimal := 2.8182
	sum := entero + int(decimal)

	fmt.Printf("%T\n", sum)

	// TEXTO:
	var greeting string = "What's going on, "
	concat := greeting + "zshell"
	inMayus := strings.ToUpper(concat)

	fmt.Printf("%T\n", inMayus)

	// BOOLEANOS:
	var isTrue bool = true

	fmt.Printf("%T\n", isTrue)

	// ARREGLOS:
	var unfixedArray [3]int = [3]int{1, 2, 3}
	fmt.Printf("%T\n", unfixedArray)

	fixedArray := [...]float64{3.14, 1.72, 1.27, 2.82}
	fmt.Printf("%T\n", fixedArray)

	// SLICES:
	variSlice := []int32{0, 1, 2, 3}
	variSlice = append(variSlice, 7)
	fmt.Printf("%T\n", variSlice)

	// Los slices permiten que se jueguen con sus elementos, añadiendo o quitandolos
	lenghtSlice := len(variSlice)
	capacitySlice := cap(variSlice)
	fmt.Println(lenghtSlice)
	fmt.Println(capacitySlice)

	// MAPAS:
	dictionary := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	fmt.Printf("%T\n", dictionary)

	// STRUCTS:
	type Perro struct {
		Nombre string
		Edad   int
	}

	persona := Perro{"Nombre": "Miel", "Edad": 8}
	fmt.Println(persona)
	fmt.Printf("%T\n", persona)
}

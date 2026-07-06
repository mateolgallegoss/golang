package main

import "fmt"

func main() {

	defer fmt.Println("END")  // La línea o bloque a la que aplica se ejecuta al final del módulo

	// CONDICIONALES:
	var edad int
	_, err := fmt.Scanln(&edad)
	if err != nil { // nil refiere a un valor nulo
		fmt.Println("Error to read:", err)
		return
	}

	if edad >= 18 {
		fmt.Println("You are an adult already!")
	} else {
		fmt.Println("You are not an adult yet!")
	}

	// Assertive | Negative programming

	if edad < 18 {
		fmt.Println("You are not an adult yet!")
		return
	}
	fmt.Println("You are an adult already!")

	// BUCLES:
	for i := 0; i < 5; i++ { // Bucle clásico (para i tal que i...)
		fmt.Printf("Iteration: %d\n", i)
	}

	n := 1
	for n <= 3 { // Bucle while (mientras n...)
		fmt.Printf("Iteration: %d\n", n)
		n++
	}

	n = 0

	for {
		n++
		if n == 7 {
			continue
		}

		fmt.Printf("n in infinitive loop: %d\n", n)
		if n >= 37 {
			break
		}
	}

	// RANGE:

	slice := []string{"one", "two", "three"}
	for index, value := range slice { // En este caso, value puede ser reemplazo, más no index, debido a que está antes de la coma
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// SWITCH:

	var value int
	_, err1 := fmt.Scanln(&value)
	if err1 != nil {
		fmt.Println("Error to read:", err1)
		return
	}

	switch value {
	case 1:
		fmt.Println("It's 1")
	case 2:
		fmt.Println("It's 2")
	default:  // Necesario para evitar errores (uso obligatorio)
		fmt.Println("It's not 1 or 2")
	}
}


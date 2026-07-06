package main

import "fmt"

// FUNCIÓN CLÁSICA:

func sum(a, b int) int { // RECOMENDADO
	return a + b
}

func product(a complex64, b complex64) complex64 { // NO RECOMENDADO
	return a * b
}

// FUNCIÓN QUE DEVUELVE MÁS DE UN VALOR:

func divider(a, b float64) (float64, float64) {
	quotient := a / b
	remainder := float64(int(a) % int(b))

	return quotient, remainder
}

func main() {

	var q, r float64
	_, err := fmt.Scanln(&q, &r)
	if err != nil {
		fmt.Println("Error to read:", err)
		return
	}
	

	qu, re := divider (q, r)

	fmt.Print("Quotient: ", qu)
	fmt.Println(", Remainder: ", re)

	
}


package main

import (
	"errors"
	"fmt"
)

// FUNCIÓN CLÁSICA:

func sum(a, b int) int { // RECOMENDADO
	return a + b
}

func product(a complex64, b complex64) complex64 { // NO RECOMENDADO
	return a * b
}

// FUNCIÓN QUE DEVUELVE MÁS DE UN VALOR:

func divider(a, b float64) (float64, error) {

	if b == 0 { // ERRORES DEVUELTOS POR MÉTODOS FRECUENTEMENTE
		return 0, errors.New("No sé dividir por 0")
	}
	quotient := a / b
	return quotient, nil
}

// FUNCIÓN QUE VARÍA EL NÚMERO DE ARGUMENTOS

func printNames(names ...string) {

	for _, name := range names {
		fmt.Println(name)
	}
}

func counter() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

func main() {

	var q, r float64
	_, err := fmt.Scanln(&q, &r)
	if err != nil {
		fmt.Println("Error to read:", err)
		return
	}

	qu, err2 := divider(q, r)

	if err2 != nil {
		fmt.Println(err2)
		return
	}

	fmt.Println("Quotient: ", qu)

	var u1, u2, u3, u4 string
	_, e := fmt.Scanln(&u1, &u2, &u3, &u4)
	if e != nil {
		fmt.Println("Error to read:", e)
		return
	}

	printNames(u1, u2, u3, u4)

	count := counter()
	fmt.Println(count)
}

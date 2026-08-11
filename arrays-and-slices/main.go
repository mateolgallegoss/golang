package main

import (
	"fmt"
	"slices"
)

func introSlices(int){
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
  return
}

func handlingSlices(int){
	var slc01 []int //La variable devuelve un nil
	slc02 := []int{} //No recomendado, inicia como un literal vacío, que puede causar pánico en compilación
	var slc1 = []int{3, 4, 1, 8}
	var slc2 = make([]int, 4, 8) //Si se realiza un append, los valores empiezan en el índice 3
	var slc3 = make([]int, 0, 8) //Si se realiza un append, los valores empiezan en el índice 0

	fmt.Println("------------------------ DEFINICIÓN Y MANEJO DE SLICES ------------------------")
	fmt.Println("DECLARACIÓN Y DEFINICIÓN DE SLICES")
	fmt.Printf("Slice 1, declarado mediante var: %d \n", slc01)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc01), cap(slc01))
	fmt.Printf("Slice 2, con declaración implícita(literal vacío): %d \n", slc02)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc02), cap(slc02))
	fmt.Printf("Slice 3, definiendo valores iniciales (longitud inicial): %d \n", slc1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc1), cap(slc1))
	fmt.Printf("Slice 4, definiendo longitud y capacidad inicial con el built-in make : %d \n", slc2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc2), cap(slc2))
	fmt.Printf("Slice 5, definiendo capacidad inicial con make, y dejando en nulo la longitud inicial: %d \n", slc3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc3), cap(slc3))
	
	fmt.Printf("¿Es igual el slice 1 al slice 2?: %b \n", slices.Equal(slc01, slc02)) //Devuelve false

	slc1 = append(slc1, 3, 7)
	slc2 = append(slc2, 3, 7)
	slc3 = append(slc3, 3, 7)

	fmt.Println("APPEND DE ELEMENTOS AL 3ER, 4TO Y 5TO SLICE")
	fmt.Printf("Slice 3: %d \n", slc1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc1), cap(slc1))
	fmt.Printf("Slice 4: %d \n", slc2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc2), cap(slc2))
	fmt.Printf("Slice 5: %d \n", slc3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc3), cap(slc3))
  return
}


func slicesOfSlice(int){
	var slc1 []string
	slc1 = append(slc1, "w", "x", "y", "z")

	var slc2 = slc1[2:] // Devuelve "y" y "z"
	var slc3 = slc1[:2] //Devuelve "w" y "x"
	var slc4 = slc1[1:3] //Devuelve "x" y "y"

	fmt.Println("------------------------ SLICES de un SLICES ------------------------")
	fmt.Println("DECLARACIÓN Y DEFINICIÓN DE SLICES")
	fmt.Printf("Slice Principal: %s \n", slc1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc1), cap(slc1))
	fmt.Printf("Slice Derivado 1: %s \n", slc2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc2), cap(slc2))
	fmt.Printf("Slice Derivado 2: %s \n", slc3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc3), cap(slc3))
	fmt.Printf("Slice Derivado 3: %s \n", slc4)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc4), cap(slc4))
  
	slc2 = append(slc2, "a")
  slc3 = append(slc3, "m", "n")
	slc4 = append(slc4, "p", "q")

	fmt.Println("APPENDS A LOS SLICES")
	fmt.Println("Se hizo apppend al Slice Derivado 1 añadiendole 'a', luego añadiendole 'm' y 'n' al segundo Derivado, \n para finalmente añadirle 'p' y 'q' al tercer Derivado")
	fmt.Printf("Slice Principal: %s \n", slc1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc1), cap(slc1))
	fmt.Printf("Slice Derivado 1: %s \n", slc2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc2), cap(slc2))
	fmt.Printf("Slice Derivado 2: %s \n", slc3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc3), cap(slc3))
	fmt.Printf("Slice Derivado 3: %s \n", slc4)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc4), cap(slc4))

	var slc5 []string
	slc5 = append(slc5, "w", "x", "y", "z")

	// El tercer valor añadido es el máximo
	// longitud = últimaPosición - 1eraPosición
	// capacidad = máximo - 1eraPosición
	var slc6 = slc5[2:4:4]
	var slc7 = slc5[:2:2]
	var slc8 = slc5[1:3:3] 

	fmt.Println("------------------------ SLICES de un SLICES, SIN COMPARTIR MEMORIA ------------------------")
	fmt.Println("DECLARACIÓN Y DEFINICIÓN DE SLICES")
	fmt.Printf("Slice Principal: %s \n", slc5)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc5), cap(slc5))
	fmt.Printf("Slice Derivado 1: %s \n", slc6)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc6), cap(slc6))
	fmt.Printf("Slice Derivado 2: %s \n", slc7)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc7), cap(slc7))
	fmt.Printf("Slice Derivado 3: %s \n", slc8)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc8), cap(slc8))
	
	slc6 = append(slc6, "a")
  slc7 = append(slc7, "m", "n")
	slc8 = append(slc8, "p", "q")

	fmt.Println("APPENDS A LOS SLICES")
	fmt.Println("Se hizo apppend al Slice Derivado 1 añadiendole 'a', luego añadiendole 'm' y 'n' al segundo Derivado, \n para finalmente añadirle 'p' y 'q' al tercer Derivado")
	fmt.Printf("Slice Principal: %s \n", slc5)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc5), cap(slc5))
	fmt.Printf("Slice Derivado 1: %s \n", slc6)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc6), cap(slc6))
	fmt.Printf("Slice Derivado 2: %s \n", slc7)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc7), cap(slc7))
	fmt.Printf("Slice Derivado 3: %s \n", slc8)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc8), cap(slc8))

	return
}

func useCopy(int){
	var slc1 = []int{1, 2, 3, 4, 5}
	var cp1 = make([]int, 3)

	fmt.Println("------------------------ FUNCIÓN COPY ------------------------")
	fmt.Println("FUNCIÓN COPY EN SLICE CON LONGITUD DEFINIDA")
	fmt.Printf("Slice a copiar: %d \n", slc1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc1), cap(slc1))
	fmt.Printf("Slice copia, solo declarado: %d \n", cp1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(cp1), cap(cp1))

	num1 := copy(cp1, slc1) // Se copian los primeros 3 elementos de slc1 a cp1 
	fmt.Printf("Slice copia definido: %d \n", cp1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n Datos copiados: %d \n", len(cp1), cap(cp1), num1)

	var slc2 = []int{1, 2, 3, 4, 5}
	var ar1 = [5]int{6, 7, 8, 9, 10}
	var cp2 = make([]int, 3)

	fmt.Println("FUNCIÓN COPY DE ARRAY A SLICE Y DE SLICE A ARRAY")
	fmt.Printf("Slice a copiar: %d \n", slc2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc2), cap(slc2))
	fmt.Printf("Array a superpuesto, con definición inicial, pero sin la superposición hecha: %d \n", ar1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(ar1), cap(ar1))
	fmt.Printf("Slice copia, solo declarado: %d \n", cp2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(cp2), cap(cp2))

	num2 := copy(cp2, ar1[:]) // Se copian los 3 primeros elemento del ar1 a cp2
	fmt.Printf("Slice copia definido: %d \n", cp2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n Datos copiados: %d \n", len(cp2), cap(cp2), num2)

	num3 := copy(ar1[:], slc2) // Se superponen los elementos de slc2 en ar1
	fmt.Printf("Array tras superponerle el slice: %d \n", ar1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n Datos copiados: %d \n", len(ar1), cap(ar1), num3)

	var slc3 = []int{1, 2, 3, 4, 5}
	var cp3 = make([]int, 3)

	fmt.Println("COPY A SLICE USANDO UN SLICE DE UN SLICE")
	fmt.Printf("Slice a copiar: %d \n", slc3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc3), cap(slc3))
	fmt.Printf("Slice copia, solo declarado: %d \n", cp3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(cp3), cap(cp3))

	num4 := copy(cp3, slc3[1:4]) // Se copian los elementos del segundo al cuarto índice 
	fmt.Printf("Slice copia definido: %d \n", cp3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n Datos copiados: %d \n", len(cp3), cap(cp3), num4)
	return
}

func transFormations(int){
	var ar1 = [4]int{1, 2, 3, 4}
	var slc1 = ar1[:]
	var slc2 = ar1[2:]
	var slc3 = ar1[:2]

	fmt.Println("------------------------ CONVERSIÓN DE ARRAYS Y SLICES ------------------------")
	fmt.Println("CONVERSIÓN DE ARRAY A SLICE")
	fmt.Printf("Array: %d \n", ar1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(ar1), cap(ar1))
	fmt.Printf("Slice con copia completo: %d \n", slc1)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc1), cap(slc1))
	fmt.Printf("Slice con copia de los dos últimos elementos: %d \n", slc2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc2), cap(slc2))
	fmt.Printf("Slice con copia de los dos primeros elementos: %d \n", slc3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc3), cap(slc3))

	ar1[1] = 37

	fmt.Println("SE REEMPLAZA EL ÍNDICE 1 DEL ARREGLO POR 37")
	fmt.Printf("Array: %d \n", ar1)
	fmt.Printf("Slice con copia completo: %d \n", slc1)
	fmt.Printf("Slice con copia de los dos últimos elementos: %d \n", slc2)
	fmt.Printf("Slice con copia de los dos primeros elementos: %d \n", slc3)

	slc2 = append(slc2, 5)
	fmt.Println("SE AÑADE AL FINAL DE SEGUNDO SLICE UN 5")
	fmt.Printf("Array: %d \n", ar1)
	fmt.Printf("Slice con copia completo: %d \n", slc1)
	fmt.Printf("Slice con copia de los dos últimos elementos: %d \n", slc2)
	fmt.Printf("Slice con copia de los dos primeros elementos: %d \n", slc3)

	var slc4 = []int{5, 6, 7, 8}
	var ar2 = [4]int(slc4) // La longitud del arreglo no puede ser mayor a la del slice, unicamente igual o menor
	var ar3 = [2]int(slc4)

	fmt.Println("CONVERSIÓN DE SLICE A ARRAY")
	fmt.Printf("Slice: %d \n", slc4)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(slc4), cap(slc4))
	fmt.Printf("Array con copia completo: %d \n", ar2)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(ar2), cap(ar2))
	fmt.Printf("Array con copia de los dos primeros elementos: %d \n", ar3)
	fmt.Printf("Longitud: %d \n Capacidad: %d \n", len(ar3), cap(ar3))
	return
}

func main(){
	introSlices(0)
	handlingSlices(0)
	slicesOfSlice(0)
	useCopy(0)
	transFormations(0) 
}

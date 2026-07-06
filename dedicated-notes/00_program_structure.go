package main // En Go, el  código se organiza en paquetes. El paquete "main" indica que el código es ejecutable y no es una biblioteca.

// Go es un lenguaje muy estricto, por lo que el código no se compila si se importan paquetes que no se van a usar.

import ( // Declara que paquetes se van a importar dentro del fichero en el que se trabaja.
	"fmt"
	"os"
)

func main() {  // Es el punto de entrada de la aplicación, usado para ejecutar el código.

	if len(os.Args) != 2 {
		os.Exit(1)
	}
    fmt.Println("Pero q mandaann", os.Args[1])
}
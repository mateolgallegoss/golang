# Structure of Go module
En Go, tenemos varias formas de organizar directorios con uno o más módulos.

## Structure of simple directory
Un directorio simple consta de un solo módulo, por lo que todos los archivos pueden ir en la ruta del directorio, sin necesidad de pertenecer a una sub-carpeta.

El directorio consta de tres archivos:

- `main.go`
  Es el fichero que se ejecuta al iniciar el programa.
- `go.mod`
  Archivo que maneja las operaciones en un módulo.
- `go.sum`
  Archivo que autentica el `go.sum`.
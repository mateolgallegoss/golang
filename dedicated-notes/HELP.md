# Un vistazo a Go, go help

En la imagen se puede previsualizar el contenido que bash devuelve tras ejecutar `go help`.

![go help en konsole KDE Plasma](media/go_help_1.png)

Las guías a mostrar a continuación se deben ir revisando conforme se desarrolle su aprendizaje en Go.

## Comandos para terminal

### go run [ruta]
Compila y ejecuta un programa ejecutable (valga la redundancia), es decir, un programa que posee el paquete main y su función respectivamente.

### go mod
Brinda acceso a operaciones en un módulo.

1. Inicializa un nuevo módulo en el directorio actual.
~~~
go mod init
~~~

2. Edita las herramientas y scripts del fichero `go.mod`.
~~~
go mod edit
~~~

3. Añade módulos faltantes y remueve los que están sin usar. Adicionalmente, crea un archivo para autentica el `go.mod`, el `go.sum`.
~~~
go mod tidy
~~~

4. Descarga módulos a la memoria caché local.
~~~
go mod download
~~~

5. Verifica que las dependencias tengan el contenido necesario (que se espera).
~~~
go mod verify
~~~

6. Explica porqué un paquete o módulo es necesario.
~~~
go mod why
~~~

7. Crea copias vendidas de las dependencias, usado para dependencias externas, que se guardan en una ruta `vendor/del_proyecto`.
~~~
go mod vendor
~~~

8. Imprime un gráfico con los requerimientos del módulo.
~~~
go mod graph
~~~

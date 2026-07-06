# Elementos léxicos

## Comentarios
Un comentario no puede venir incluido dentro de un string literal o una runa (characteres unicode).
Cualquier comentario que pertenece a una línea con código, se lo considera un espacio; en cambio, a cualquier otro se lo considera una nueva línea.

**Formas de uso:**
1. Usando `//` para generar un comentario en una sola línea.
2. Utilizando `/*` para abrir un comentario multilíneas, y cerrando con su subsequente `*/`.

## Tokens
Existen 4 clases de tokens en Go:
- Identificadores
- Keywords (Palabras claves)
- Operadores y puntuadores
- Literales
Los espacios blancos formados por espacios, tab horizontal, carriage return y líneas nuevas, son ignorados a no ser que separen tokens, lo que causa que se combinen en un solo token.

### Punto y coma, tokens
Además, una línea nueva o el final del fichero son tomados como sí se insertara un `punto y coma`. Tener en cuenta que al tokenizar, el `lexer` (analizador léxico) toma la secuencia más larga de caracteres que forman un token valido.
En resumen, en GO, el `lexer` automáticamente añade un punto y coma al final de cada token.

**NOTA**: *Los tokens son el bloque base que permite al interpreté de Go compilar todo el código y traducirlo.

## Punto y coma

El semicolon, como se lo nombra en inglés, usa el símbolo de punto y coma `;`.
Los programas siguen dos reglas paraa omitir la mayoría de puntos y comas en la medida de lo posible:
1. La entrada se fragmenta en tokens, insertando un punto y coma en el token al final del mismo, siempre y cuando el token sea del tipo:
  - Un identificador  
  - Entero, flotante, complejo, rune o un string literal
  - Una de las keywords: `break`, `continue`, `falltrought` o `return`.
  - Uno de los operadores o puntuadores: `++`, `--`, `)`, `]` o `}`.

2. Al usar bloques complejos en una sola línea, el punto y coma se omite previo al cierre de un `)` o `}`.

## Identificadores

Los identificadores sirven para nombrar entidades del programa, como variables y tipos, siedo una secuencia de una o más letras y números. El nombre de un identificador siempre debe comenzar con una letra.
~~~
z
_x37
αβ
camelCase
AguanteBackEnd
~~~

## Keywords

Las siguientes palabras claves están reservadas y no se pueden usar como identificadores:
~~~
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
~~~

## Operadores
Los siguientes símbolos son reservados a operadores, operadores asignados y puntuadores:

### Operadores
~~~
+ - * / % & | 
^ << >> &^ && 
|| <- ++ -- == !=
< <= > >= ! ~
~~~

### Operadores asignados
~~~
= += -= *= /=
%= &= |= ^=
<<= >>= &^= :=
~~~

### Puntuadores
~~~ 
( ) [ ] { } ,
; . : ...
~~~


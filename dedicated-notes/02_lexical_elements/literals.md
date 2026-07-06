# Literales

**NOTA:** por criterío de eficiencia lectora, una barra baja debe aparecer después de un prefijo base o entre una secuencia de dígitos.

## Literales de tipo entero
Son literales que representan una constante entera con una secuencia de dígitos, existen los siguientes literales de tipo entero:
- Binario: Representados con 0b o 0B.
- Octonario: 0, 0o y 0O son usados para este literal.
- Décimal: Un número sin acompañar es simplemente un número decimal
- Hexadécimal: Usa números y letras. a/A hasta f/F representan los números del 10 al 15.

~~~
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .
~~~

## Literales de tipo punto flotante
Es un literal que representan con dígitos decimales o hexadecimales una constante de punto flotante.

EL literal con décimal consite de una parte entera (decimal), un punto decimal, una parte fraccionaria y la parte exponencial (decimal), que lleva un `e/E` seguida del signo (sí es negativo) y digitos

El literal con hexadecimal tiene un prefijo `0x/0X`, una parte entera (hexadecimal), un punto base, una parte fraccionaria (hexadécimal) y la parte exponencial (decimal), que es una `p/P` seguida del signo (sí es negativo) y dígitos.

Se puede prescindir del punto y la parte fraccionaria, más no del exponente.

~~~
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
~~~

## Literales de tipo imaginario
Este literal representa la parte imaginaria de una constante compleja. Se construye usando un literal de tipo entero o punto flotante, seguido de la letra "i" en `lower_case`.

EL valor del literal se obtiene del producto entre el literal de tipo entero o punto flotante y la unidad imaginaria "i".

~~~
imaginary_lit = (decimal_digits | int_lit | float_lit) "i" .
~~~

## Literal de tipo runa
Un literal de tipo runa representa una constante rúnica en un punto del código Unicode. Es representado con unos o más caracteres englosados entre comillas, que pueden tener la forma "x" o "/a".

Un simple carácter entre comillas representa su valor Unicode, mientras que varios caracteres precedidos por una barra invertida, encodifican valores en varios formatos.



~~~
rune_lit         = "'" ( unicode_value | byte_value ) "'" .
unicode_value    = unicode_char | little_u_value | big_u_value | escaped_char .
byte_value       = octal_byte_value | hex_byte_value .
octal_byte_value = `\` octal_digit octal_digit octal_digit .
hex_byte_value   = `\` "x" hex_digit hex_digit .
little_u_value   = `\` "u" hex_digit hex_digit hex_digit hex_digit .
big_u_value      = `\` "U" hex_digit hex_digit hex_digit hex_digit
                           hex_digit hex_digit hex_digit hex_digit .
escaped_char     = `\` ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | `\` | "'" | `"` ) .
~~~

## Literal de tipo string
Este literal representa una constante string, obtenida al concatenar varios caractéres. Existen dos tipos de literales de cadena:

### Literales cadena sin formato / inenterpretada
Son literales de cadena que van entre comillas invertidas `'foo'`, representado carácteres inenterpretados (UTF-8 encodificado implícito). Las barras invertidas `/` y el carriage return `'\r'` no tienen significado en el valor de la cadena inenterpretada.

### Literales cadena interpretados
Son literales que van entre comillas dobles `"bar"`, e interpreta cualquier valor dentro de las comillas excepto nuevas líneas o comillas dobles. Este literal si incluye una barra invertida es interpretado como un literal de runa. El octal de tres dígitos y el hexadecimal de dos dígitos son representados con `bytes` individuales de la cadena resultante.
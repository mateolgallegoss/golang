# Constantes

Las constantes son representadas por un literal entero, de punto flotante, complejo, rúnico o cadena.
Las constantes numéricas representan un valor arbitrario y no uno desbordante, por ello, los literales de cadena no son parte de este grupo.

## Constantes tipadas
Una constantes es tipada cuando se de manera explicita se hace una declaración de constante o una conversión, mientras que de manera explícita, cuando se hace una declaración de variable, se le asigna un valor a una variable ya declarada, o se la usa para operar. En caso de asigar o interpretar con un tipo diferente al de la constante, se genera un error que devuelve una constante `No-typed`.

## Constantes no tipadas
Una constante no tipada obtienen un valor por defecto cuando es convertida a un tipo en base al contexto de uso, como por ejemplo, sucede en una declaración corta de variable (`short variable declaration`). El tipo por defecto de una constante no tipada puede ser un booleano, entero, rúnico, flotante64, complejo128 o una cadena.

### NOTA:
Las constantes numérica que son arbitrarias en el lenguaje, hacen que el `compilador` realize una representación interna limitada, por lo que se debe tener en cuenta lo siguiente:
- **Representa el valor entero de una constante con al menos 256 bits.**
- **Representa un punto flotante, incluyendo la parte de una constante compleja, con una mantisa de al menos 256 bits y un exponente binario de al menos 16 bits.**
- **Retorna un error si no se puede representar precisamente una constante entera.**
- **Retorna un error si no se puede representar una constante de punto flotante o compleja debido al desbordamiento.**
- **Redondea a la representación de constante más cercana si no se puede representar una constante de punto flotante o compleja por los límites de precisión**
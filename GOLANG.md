Google. Robert Griesemer, Rob Pike y Ken Thompson desarrollo 2007 open source 2009


Inspirado de Algol, Pascal, C, Modula, Oberon, Smalltalk and Newsqueak y de un paper sobre procesos de comunicación secuencial

Golang es un lenguaje de programación de código abierto, de propósito general y soporta múltiples paradigmas.

No es un lenguaje orientado a objetos y sin embargo nos permite implementar los conceptos que engloban POO:

Abstracción
Encapsulamiento
Polimorfismo
Herencia
Modularidad
Principio de ocultación
Recolección de basura

Go Way--> Go idiomático

No existen excepciones, sí manejo de errores

Alternativa a herencia-->Composición

Tipos 

abstractos

*definen scopes para los tipos concretos
*definen comportamientos sin decidir un receiver
*no se sabe el tipo en tiempo de compilación
*menos eficiente
*fácil de interceptar


concretos

*layout en memoria
*se definen comportamientos en base a su tipo de dato a través de métodos
*conocidos en compilación
*muy eficientes
*no se pueden interceptar




Mayor implementación de composición es a través de interfaces

Las interfaces se pueden unir, pero ojo! mientras más abarca una interfaz más débil es la abstracción
La abstracción se acaba cuando aparece un tipo concreto

Las interfaces rompen las dependencias, desacoplando 

Más fácil de testear, más fácil de reutilizar

Una interfaz define un comportamiento de un tipo, contiene una colección de firmas de métodos

Para implementar una intefaz en Go simplemente implementamos todos sus métodos

Si una variable es del tipo de una interfaz entonces puede llamar a todos sus métodos

A su vez, también nos sirve para aplicar polimorfismo.

No existe el implements de java(implementacion explicita)

En Go existe de forma implicita, podemos definir que una función pertenece a una estructura a través de Receiver functions

Las interfaces proveen 

algoritmos genéricos
esconde las implementaciones
métodos de intercepción

podemos usar type Assertion

- entender comportamientos
- clasificar errores
- mantener compatibilidad

Los métodos en una interfaz pueden ir con parámetros y especificando el retorno o solo describiendo el nombre de la función(depende de cómo querramos definir la firma)


Repositorio

go get -u github.com/go-sql-driver/mysql
go get github.com/DATA-DOG/go-sqlmock


Links

https://golang.org/doc/faq#Is_Go_an_object-oriented_language
https://blog.friendsofgo.tech/posts/es_go_un_lenguaje_orientado_a_objetos/
https://blog.learngoprogramming.com/about-go-language-an-overview-f0bee143597c
https://www.youtube.com/watch?v=KINIAgRpkDA
https://manifold.co/blog/how-go-interfaces-can-facilitate-switching-external-services-619cc478e20a
https://www.integralist.co.uk/posts/go-interfaces/
https://youtu.be/F4wUrj6pmSI

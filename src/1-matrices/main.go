package main

import "fmt"




func main() {
	//! declaracion de MATRICES
	//var matriz [3]int //en los corchetes se pone el tamaño de la matriz //SE INICIALIZA EN 0 POR DEFECTO -> [0 0 0]

	var a [5]int
	a[2] = 7 // [0 0 7 0 0]

	//inicializo con valores
	var b = [5]int{1,2,3,4,5} // [1 2 3 4 5]
	b[2] *= 10 // [1 2 30 4 5]

	//cuando no se la cantidad de datos que va a tener la matriz
	c := [...]int{1,2,3} // [1 2 3]
	
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
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

	//recorrer una matriz con FOR
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}

	//recorrer una matriz con RANGE //index, value
	for i, v := range c {
		//fmt.Println(i, v) //devuelve el index y el valor
		fmt.Printf("indice: %d, valor: %d\n", i, v)
	}

	//! matrices BI-DIMENSIONALES

	var matrix = [3][2]int{
		{1,2},
		{3,4},
		{5,6},
	}

	fmt.Println(matrix) // [[1 2] [3 4] [5 6]]


	
}
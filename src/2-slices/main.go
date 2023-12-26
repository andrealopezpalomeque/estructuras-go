package main

import "fmt"

func main(){
	//! declarar un slice
	//var a []int 
	//! agrego elementos al slice usando append
	//slice = append(slice, elemento)
	//a = append(a, 1)

	//otra forma de declarar
	diasSemana := []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}


	//! crear un slice a partir de otro slice
	//slice[desde:hasta]
	//desde es inclusivo, hasta es exclusivo [)
	//fmt.Println(diasSemana[0:3])
	//! longitud del slice
	fmt.Println(len(diasSemana))
	//! capacidad del slice
	fmt.Println(cap(diasSemana)) //capacidad es el DOBLE de la longitud

	//! eliminar elementos de un slice
	//slice = append(slice[:indice], slice[indice+1:]...)
	diasSemana = append(diasSemana[:3], diasSemana[4:]...) //elimino el elemento en el indice 3 (Jueves) y sigo con el resto desde el indice 4 (Viernes)
	fmt.Println(diasSemana)

	
	
	//--------------------------------------------//
	// !Funcion make
	//make(tipo, longitud, capacidad)
	nombres := make([]string, 5, 10)
	nombres[0] = "Juan"
	fmt.Println(nombres)

	// !Funcion copy
	//copy(destino, fuente)
	slice1 := []int{1,2,3,4,5}
	slice2 := make([]int, 5)
	copy(slice2, slice1) //copio el slice1 en el slice2, si pongo al reves no tiene valores para copiar
	fmt.Println(slice1)
	fmt.Println(slice2)


	

	
	
}
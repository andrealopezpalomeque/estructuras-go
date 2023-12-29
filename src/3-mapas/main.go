package main

import "fmt"


func main(){

	//! declarar un mapa
	//clave:valor

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	//! agrego un nuevo color
	colors["yellow"] = "#lakjsdf"

	fmt.Println(colors["white"]) // #ffffff -> le paso la CLAVE y me devuelve el VALOR
	fmt.Println(colors) // map[green:#4bf745 red:#ff0000 white:#ffffff yellow:#lakjsdf]

	//! guardo el valor de un color en una variable
	//! agrego verificacion de si existe el color con ok
	white, ok := colors["white"]

	if ok {
		fmt.Println(white) // #ffffff
	} else {
		fmt.Println("No existe la clave")
	}

	//! eliminar un elemento
	//mapa, clave a eliminar
	delete(colors, "white")

	fmt.Println(colors) // map[green:#4bf745 red:#ff0000 yellow:#lakjsdf]

	//! iterar un mapa
	for clave, valor := range colors {
		fmt.Println("CLAVE:", clave, "VALOR:", valor)
	}


}
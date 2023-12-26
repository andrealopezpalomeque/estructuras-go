package main

import "fmt"

type Persona struct {
	nombre string
	edad int
	correo string
}


func main(){
	var p Persona

	p.nombre = "Juan"
	p.edad = 20
	p.correo = "juan@g.com"

	fmt.Println(p) // {Juan 20 andrea@g}

	//! se puede hacer en una sola linea
	p2 := Persona{"Andrea",25,"andrea@g.com"}

	fmt.Println(p2) // {Andrea 30 juan@g}
	

}
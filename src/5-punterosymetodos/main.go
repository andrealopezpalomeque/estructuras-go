package main

import "fmt"

type Persona struct{
	nombre string
	edad int
	correo string
}

//! UN METODO PERTECE A UNA INSTANCIA DE UNA ESTRUCTURA
func (p *Persona) saludar(){
	fmt.Println("Hola mi nombre es", p.nombre)
}


func main(){
	var x int = 10
	var p *int = &x // p es un puntero a int que apunta a la dirección de memoria de x

	fmt.Println(x) // 10
	fmt.Println(p) // 0xc0000a4008

	editar(&x)
	fmt.Println(x) // 20

	// ! PERSONA
	p1 := Persona{"Juan", 20, "juan@g.com"}
	p1.saludar() //accedo al metodo mediante la instacia de la persona 
	
	
	
	
}

//! LAS FUNCIONES SON INDEPENDIENTES
func editar(x *int){
	*x = 20
}
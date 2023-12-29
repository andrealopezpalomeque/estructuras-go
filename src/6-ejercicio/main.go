package main

import (
	"bufio"
	"fmt"
	"os"
)

//---------------ESTRUCTURAS-----------------//
type Tarea struct {
	nombre string
	desc string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea //slice que almacena un tipo de dato Tarea
}

//-------------------METODOS-------------------//S

//metodo AGREGAR TAREA
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t) //append agrega un elemento al slice 
}

//metodo MARCAR COMO COMPLETADA
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true //busco la tarea por el indice y modifico el campo COMPLETADO
}

//metodo EDITAR TAREA
func (l *ListaTareas) editarTarea(index int, nombre string) {
	l.tareas[index].nombre = nombre //busco la tarea por el indice y modifico el campo NOMBRE
}

//metodo ELIMINAR TAREA
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...) //elimino la tarea por el indice
}

func main(){
	//creo una lista de tareas
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)


	for{
		var opcion int
		fmt.Println("Seleccione una opcion:\n",
			"1. Agregar tarea\n",
			"2. Marcar como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir",
			)
		fmt.Print("Ingrese la opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre , _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.desc , _ = leer.ReadString('\n')
			t.completado = false
			lista.agregarTarea(t)
		case 2:
			var index int
			fmt.Print("Ingrese el numero de la tarea a marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el numero de la tarea a editar: ")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nuevo nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			lista.editarTarea(index, t.nombre)
		case 4:
			var index int
			fmt.Print("Ingrese el numero de la tarea a eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
		case 5:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opcion no valida")

		}

		fmt.Println("Lista de tareas:")
		fmt.Println("--------------------------------------------------")
		//listar TODAS LAS TAREAS
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
			
		}
		fmt.Println("--------------------------------------------------")
	}


}
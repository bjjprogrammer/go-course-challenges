package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name        string
	description string
	done        bool
}

type TaskList struct {
	tasks []Task
}

func (t *TaskList) addTask(task Task) {
	t.tasks = append(t.tasks, task)
}

func (t *TaskList) completeTask(index int) {
	t.tasks[index].done = true
}

func (t *TaskList) editTask(index int, task Task) {
	t.tasks[index] = task
}

func (t *TaskList) removeTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func main() {
	list := TaskList{}

	read := bufio.NewReader(os.Stdin)
	for {
		var option int
		fmt.Println("Seleccione una opción: \n 1. Agregar tarea \n 2. Completar tarea \n 3. Editar tarea \n 4. Eliminar tarea \n 5. Salir")
		fmt.Println("Ingrese una opción:")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var t = Task{}
			fmt.Println("Ingrese el nombre de la tarea:")
			t.name, _ = read.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			t.description, _ = read.ReadString('\n')

			list.addTask(t)
		case 2:
			var index int
			fmt.Println("Ingrese el número de la tarea que desea completar:")
			fmt.Scanln(&index)
			list.completeTask(index)
			fmt.Println("Tarea completada")
		case 3:
			var index int
			var t = Task{}
			fmt.Println("Ingrese el número de la tarea que desea editar:")
			fmt.Scanln(&index)
			fmt.Println("Ingrese el nombre de la tarea:")
			t.name, _ = read.ReadString('\n')
			fmt.Println("Ingrese la descripción de la tarea:")
			t.description, _ = read.ReadString('\n')
			list.editTask(index, t)
		case 4:
			var index int
			fmt.Println("Ingrese el número de la tarea que desea eliminar:")
			fmt.Scanln(&index)
			list.removeTask(index)
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida")
		}

		fmt.Println("Task List:")
		fmt.Println("==========")

		for i, task := range list.tasks {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, task.name, task.description, task.done)
		}

		fmt.Println("==========")
	}

}

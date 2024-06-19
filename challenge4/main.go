package main

import "library/animals"

func main() {
	// Create a slice of animals
	data := []animals.Animal{
		animals.NewDog("Fido"),
		animals.NewCat("Whiskers"),
	}

	// Show the behavior of each animal
	animals.ShowAnimalBehavior(data)
}

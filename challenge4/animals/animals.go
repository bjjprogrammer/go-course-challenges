package animals

import "fmt"

type Animal interface {
	Speak() string
	Move() string
}

type Dog struct {
	Name string
}

func NewDog(name string) *Dog {
	return &Dog{Name: name}
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Walk"
}

type Cat struct {
	Name string
}

func NewCat(name string) *Cat {
	return &Cat{Name: name}
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Walk"
}

func ShowAnimalBehavior(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.Speak())
		fmt.Println(animal.Move())
	}
}

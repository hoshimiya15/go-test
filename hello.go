package main

import (
	"fmt"
)

type IAnimal interface {
	GetName() string
}

type Dog struct {
	name string
}

func (d Dog) GetName() string {
	return d.name
}

type Cat struct {
	name string
}

func (c Cat) GetName() string {
	return c.name
}

func main() {
	var animals []IAnimal = []IAnimal{
		Dog{"ポチ"},
		Dog{"ジョン"},
		Cat{"ミケ"},
		Cat{"タマ"},
	}

	for _, animal := range animals {
		fmt.Println(animal.GetName())
	}
}	

package factory

import (
	"github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang_factory/animal"
	"github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang_factory/food"
)

type Factory struct {
	a animal.InterfaceAnimal
	f food.InterfaceFood
}

func (f Factory) AnimalFactoryMethod() animal.InterfaceAnimal {
	return f.a
}

func (f Factory) FoodFactoryMethod() food.InterfaceFood {
	return f.f
}

func NewFactory(a animal.InterfaceAnimal, f food.InterfaceFood) Factory {
	var ai animal.InterfaceAnimal
	var fi food.InterfaceFood

	switch a.(type) {
	case animal.Dog:
		ai = animal.NewDog()
	case animal.Cat:
		ai = animal.NewCat()
	default:
		panic("error")
	}

	switch f.(type) {
	case food.Meat:
		fi = food.NewMeat(ai)
	case food.Fish:
		fi = food.NewFish(ai)
	default:
		panic("error")
	}
	return Factory{ai, fi}
}

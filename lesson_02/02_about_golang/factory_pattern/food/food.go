package food

import (
	"strconv"
	"github.com/akeyace/learning-golang-over-OOP/lesson_02/02_about_golang/factory_pattern/animal"
)

type (
	InterfaceFood interface {
		Get() string
	}

	Food struct {
		Volume int
	}

	Meat struct {
		Food
	}

	Fish struct {
		Food
	}

	Mud struct {
		Food
	}
)

func NewMeat(a animal.InterfaceAnimal) Meat {
	switch a.(type) {
	case animal.Dog:
		return Meat{Food{100}}
	case animal.Cat:
		return Meat{Food{40}}
	default:
		panic("error")
	}
}

func NewFish(a animal.InterfaceAnimal) Fish {
	switch a.(type) {
	case animal.Dog:
		return Fish{Food{100}}
	case animal.Cat:
		return Fish{Food{40}}
	default:
		panic("error")
	}
}

func (m Meat) Get() string {
	return "this is a meat! volume: " + strconv.Itoa(m.Volume) + "g"
}

func (f Fish) Get() string {
	return "this is a fish! volume: " + strconv.Itoa(f.Volume) + "g"
}

func (m Mud) Get() string {
	return "this is a mud.... volume: " + strconv.Itoa(m.Volume) + "g"
}

package animal

type (
	InterfaceAnimal interface {
		GetName() string
		Say() string
	}

	Animal struct {
		Name string
	}

	Dog struct {
		Animal
	}

	Cat struct {
		Animal
	}

	Monster struct {
		Animal
	}
)

func (a Animal) GetName() string {
	return "I'm " + a.Name
}

func (d Dog) Say() string {
	return "ワン"
}

func (c Cat) Say() string {
	return "ニャン"
}

func (m Monster) Say() string {
	return "ぎゃー"
}

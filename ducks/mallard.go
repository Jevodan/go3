package ducks

import (
	"ducks2/fly"
	"fmt"
)

var mallard = mallardDuck{}

type mallardDuck struct {
	Duck
}

func NewMallardDuck() mallardDuck {
	mallard.name = "Утка мандаринка"
	mallard.flyBehavior = &fly.FlyWithWings{}
	//	mallard.quackBehavior = &Squeak{}
	return mallard
}

func (m *mallardDuck) Display() {
	fmt.Println("Утка маллард дак на экране!!!!", m.GetName())
}

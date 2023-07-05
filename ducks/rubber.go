package ducks

import (
	"ducks2/fly"
	"fmt"
)

var rubber = rubberDuck{}

type rubberDuck struct {
	Duck
}

func NewRubberDuck() *rubberDuck {
	rubber.name = "Утка манок"
	rubber.flyBehavior = &fly.NoFly{}
	return &rubber
}

func (r *rubberDuck) Display() {
	fmt.Println(r.GetName(), " на экране!!!!")
}

/*
func (r *RubberDuck) Quack() {
	fmt.Println("Резиновая утка", r.GetName(), " пищит!!!!")
}
*/

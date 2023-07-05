package ducks

import "fmt"

var red = redDuck{}

type redDuck struct {
	Duck
}

func (r *redDuck) Display() {
	fmt.Println(r.GetName(), " на экране!!!!")
}

func NewRedDuck() *redDuck {
	red.name = "Красная утка"
	//	mallard.flyBehavior = &FlyWithWings{}
	//	mallard.quackBehavior = &Squeak{}
	return &red
}

/*
func (r *RedDuck) Quack() {
	fmt.Println("Утка ", r.GetName(), " крякает")
}

func (r *RedDuck) Fly() {
	fmt.Println("Утка ", r.GetName(), " летает")
}
*/

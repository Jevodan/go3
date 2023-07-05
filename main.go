package main

import (
	"ducks2/ducks"
)

func main() {

	mallard := ducks.NewMallardDuck()
	red := ducks.NewRedDuck()
	rubber := ducks.NewRubberDuck()

	//mallard.SetFlyBehavior(&ducks.NoFly{})

	var ducksArr []ducks.Duck
	ducksArr = append(ducksArr, red.Duck, mallard.Duck, rubber.Duck)

	for _, value := range ducksArr {
		value.PerformFly()
		value.Swim()
		//	value.PerformQuack()
		value.Display()
	}
}

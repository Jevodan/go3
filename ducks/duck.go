package ducks

import (
	"ducks2/fly"
	"fmt"
)

type Duck struct {
	name        string
	flyBehavior fly.FlyBehavior
}
func (d *Duck) Swim() {
	fmt.Println("Утка ", d.GetName(), "  плавает")
}

func (d *Duck) Display() {
	fmt.Println("Утка ", d.GetName(), "  перед нами")
}

func (d *Duck) GetName() string {
	return d.name
}

func (d *Duck) PerformFly() {
	d.flyBehavior.Fly(d.name)
}

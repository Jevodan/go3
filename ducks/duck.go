package ducks

import "fmt"

type Duck struct {
	name string
}

func (d *Duck) Display() {
	fmt.Println("Утка ", d.GetName(), "  перед нами")
}

func (d *Duck) GetName() string {
	return d.name
}

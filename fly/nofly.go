package fly

type NoFly struct {
	FlyBehavior
}

func (f *NoFly) Fly(name string) {

}

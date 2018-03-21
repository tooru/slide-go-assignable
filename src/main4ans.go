package main

import "fmt"

type Talker interface {
	Talk()
}
type Greeter struct {
	name string
}
type Greeter2 struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, %s\n", g.name)
}

func (g *Greeter2) Talk() {
	fmt.Printf("Hello, %s!\n", g.name)
}
func main() {
	var talker Talker
	talker = Greeter{"wozozo"}   // (1) OK
	talker = &Greeter{"wozozo"}  // (2) OK
	talker = Greeter2{"wozozo"}  // (3) NG method set of type Greeter2 is empty.
	talker = &Greeter2{"wozozo"} // (4) OK

	var greeter Greeter
	greeter = Greeter{"wozozo"}  // (5) OK
	greeter = &Greeter{"wozozo"} // (6) NG Greeter and *Greeter are different types.

	var greeterPtr *Greeter
	greeterPtr = Greeter{"wozozo"}  // (7) NG *Greeter and Greeter are different types.
	greeterPtr = &Greeter{"wozozo"} // (8) OK

	var talker2 Talker
	talker2 = &talker // (9) NG method set of &talker (type *Talker) is empty.

	var talkerPtr *Talker
	talkerPtr = talker // (10) NG *Talker and Talker are different types.

	talker.Talk()
	greeter.Talk()
	greeterPtr.Talk()
}

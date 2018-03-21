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

// END OMIT
func main() {
	var talker Talker
	talker = Greeter{"wozozo"}   // (1)
	talker = &Greeter{"wozozo"}  // (2)
	talker = Greeter2{"wozozo"}  // (3)
	talker = &Greeter2{"wozozo"} // (4)

	var greeter Greeter
	greeter = Greeter{"wozozo"}  // (5)
	greeter = &Greeter{"wozozo"} // (6)

	var greeterPtr *Greeter
	greeterPtr = Greeter{"wozozo"}  // (7)
	greeterPtr = &Greeter{"wozozo"} // (8)

	var talker2 Talker
	talker2 = &talker // (9)

	var talkerPtr *Talker
	talkerPtr = talker // (10)

	talker.Talk()
	greeter.Talk()
	greeterPtr.Talk()
}

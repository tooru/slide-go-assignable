package main

import "fmt"

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, %s\n", g.name)
}
func main() {
	var greeter Greeter
	var greeterPtr *Greeter

	greeter = Greeter{"wozozo"}  // (1)
	greeter = &Greeter{"wozozo"} // (2)

	greeterPtr = Greeter{"wozozo"}  // (3)
	greeterPtr = &Greeter{"wozozo"} // (4)

	greeter.Talk()
	greeterPtr.Talk()
}

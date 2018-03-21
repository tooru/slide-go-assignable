package main

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
}
func main() {
	var greeter Greeter
	var greeterPtr *Greeter

	greeter = Greeter{"wozozo"}     // (1)
	greeter = &Greeter{"wozozo"}    // (2)
	greeterPtr = Greeter{"wozozo"}  // (3)
	greeterPtr = &Greeter{"wozozo"} // (4)

	greeter.Talk()
	greeterPtr.Talk()
}

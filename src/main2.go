package main

import "fmt"

type Talker interface {
	Talk()
}
type Greeter struct {
	name string
}

func (g *Greeter) Talk() { // レシーバは*Greeter型、、 // HL
	fmt.Printf("Hello, %s\n", g.name)
}
func main() {
	var talker Talker
	talker = Greeter{"wozozo"} // Greeter型は代入できる？？ // HL
	talker.Talk()
}

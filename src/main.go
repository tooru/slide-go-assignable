package main

import "fmt"

type Talker interface {
	Talk()
}
type Greeter struct {
	name string
}

func (g Greeter) Talk() { // レシーバがGreeter型だが、、 // HL
	fmt.Printf("Hello, %s\n", g.name)
}
func main() {
	var talker Talker
	talker = &Greeter{"wozozo"} // Greeterのポインタ型が代入できる！なぜ？？ // HL
	talker.Talk()
}

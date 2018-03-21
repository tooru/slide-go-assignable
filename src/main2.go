package main

import "fmt"

type Talker interface {
	Talk()
}
type Greeter struct {
	name string
}

func (g *Greeter) Talk() { // レシーバ型をGreeter -> *Greeterに変更 // HL
	fmt.Printf("Hello, %s\n", g.name)
}
func main() {
	var talker Talker
	talker = Greeter{"wozozo"} // リテラルの型を*Greeter -> Greeterに変更 // HL
	talker.Talk()
}

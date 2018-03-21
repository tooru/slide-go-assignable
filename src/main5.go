package main

import "fmt"

type Greeter struct {
	name string
}

func (g Greeter) Talk() {
	fmt.Printf("Hello, %s\n", g.name)
}

func (g *Greeter) Talk2() {
	fmt.Printf("Hello, %s!\n", g.name)
}

// END OMIT
func main() {
	var talker = Greeter{"wozozo"}
	var talkerPtr = &Greeter{"wozozozo"}

	talker.Talk()                // (1)
	talker.Talk2()               // (2)
	talkerPtr.Talk()             // (3)
	talkerPtr.Talk2()            // (4)
	(&talker).Talk()             // (5)
	(&talker).Talk2()            // (6)
	(*talkerPtr).Talk()          // (7)
	(*talkerPtr).Talk2()         // (8)
	Greeter{"wozozo"}.Talk()     // (9)
	Greeter{"wozozo"}.Talk2()    // (10)
	(&Greeter{"wozozo"}).Talk()  // (11)
	(&Greeter{"wozozo"}).Talk2() // (12)
}

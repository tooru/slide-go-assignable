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

	talker.Talk()                // (1) OK
	talker.Talk2()               // (2) OK same as (&talker).Talk2()
	talkerPtr.Talk()             // (3) OK same as (*talkerPtr).Talk()
	talkerPtr.Talk2()            // (4) OK
	(&talker).Talk()             // (5) OK same as (*(&talkerPtr)).Talk()
	(&talker).Talk2()            // (6) OK
	(*talkerPtr).Talk()          // (7) OK
	(*talkerPtr).Talk2()         // (8) OK same as (&(*talkerPtr)).Talk()
	Greeter{"wozozo"}.Talk()     // (9) OK
	Greeter{"wozozo"}.Talk2()    // (10) NG cannot take the address of Greeter literal
	(&Greeter{"wozozo"}).Talk()  // (11) OK
	(&Greeter{"wozozo"}).Talk2() // (12) OK
}

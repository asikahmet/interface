package main

import "fmt"

type S struct{ v string }

func (p S) V() string { return p.v }

type Sp struct{ v string }

func (p *Sp) V() string { return p.v }

type I interface {
	V() string
}

func main() {
	s := I(S{v: "s"})
	fmt.Println(s.V())

	sp := I(&Sp{"sp"})
	fmt.Println(sp.V())
}

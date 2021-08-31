package main

import "fmt"

type Human struct {
	name string
}

type Action struct {
	Human
}

func (a Action) Move() {
	fmt.Println(a.Human.name)
}

func main() {
	a := Action{Human: Human{name: "John"}}

	fmt.Println(a.name)
}

package main

import c "patterns/chain_of_responsibility"

func main() {
	handlers := &c.ConcreteHandlerA{
		Next: &c.ConcreteHandlerB{
			Next: &c.ConcreteHandlerC{},
		},
	}
}
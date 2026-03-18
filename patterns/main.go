package main

import c "patterns/chain_of_responsibility"

func main() {
	handlers := &c.ConcreteHandlerA{
		next: &c.ConcreteHandlerB{
			next: &c.ConcreteHandlerC{},
		},
	}


}